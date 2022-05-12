package auth

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/bhojpur/data/pkg/internal/errors"

	oidc "github.com/coreos/go-oidc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	// ContextTokenKey is the key of the auth token in an
	// authenticated context
	ContextTokenKey = "authn-token"

	// The following constants are Subject prefixes. These are prepended to
	// subject names in ACLs, group membership, and any other references to subjects
	// to indicate what type of Subject or Principal they are (every Bhojpur Data
	// Subject has a logical Principal with the same name).

	// UserPrefix indicates that this Subject is a Bhojpur Data user synced from an IDP.
	UserPrefix = "user:"

	// RobotPrefix indicates that this Subject is a Bhojpur Data robot user. Any
	// string (with this prefix) is a logical Bhojpur Data robot user.
	RobotPrefix = "robot:"

	// InternalPrefix indicates that this Subject is internal to Bhojpur Data itself,
	// created to run a background task
	InternalPrefix = "internal:"

	// PipelinePrefix indicates that this Subject is a PPS pipeline. Any string
	// (with this prefix) is a logical PPS pipeline (even though the pipeline may
	// not exist).
	PipelinePrefix = "pipeline:"

	// DataPrefix indicates that this Subject is an internal Bhojpur Data user.
	DataPrefix = "data:"

	// GroupPrefix indicates that this Subject is a group.
	GroupPrefix = "group:"

	// RootUser is the user created when auth is initialized. Only one token
	// can be created for this user (during auth activation) and they cannot
	// be removed from the set of cluster super-admins.
	RootUser = "data:root"

	// ClusterAdminRole is the role for cluster admins, who have full access to all APIs
	ClusterAdminRole = "clusterAdmin"

	// RepoOwnerRole is a role which grants access to read, write and modify the role bindings for a repo
	RepoOwnerRole = "repoOwner"

	// RepoWriterRole is a role which grants ability to both read from and write to a repo
	RepoWriterRole = "repoWriter"

	// RepoReaderRole is a role which grants ability to both read from a repo
	RepoReaderRole = "repoReader"

	// IDPAdminRole is a role which grants the ability to configure OIDC apps.
	OIDCAppAdminRole = "oidcAppAdmin"

	// IDPAdminRole is a role which grants the ability to configure identity providers.
	IDPAdminRole = "idpAdmin"

	// IdentityAdmin is a role which grants the ability to configure the identity service.
	IdentityAdminRole = "identityAdmin"

	// DebuggerRole is a role which grants the ability to produce debug dumps.
	DebuggerRole = "debugger"

	// RobotUserRole is a role which grants the ability to generate tokens for robot
	// users.
	RobotUserRole = "robotUser"

	// LicenseAdminRole is a role which grants the ability to register new
	// datasvrs with the license server, manage datasvrs and update the enterprise license.
	LicenseAdminRole = "licenseAdmin"

	// AllClusterUsersSubject is a subject which applies a role binding to all authenticated users
	AllClusterUsersSubject = "allClusterUsers"

	// SecretAdminRole is a role which grants the ability to manage secrets
	SecretAdminRole = "secretAdmin"

	// DataSvrLogReaderRole is a role which grants the ability to pull datasvr logs
	DataSvrLogReaderRole = "datasvrLogReader"
)

var (
	// ErrNotActivated is returned by an Auth API if the Auth service
	// has not been activated.
	//
	// Note: This error message string is matched in the UI. If edited,
	// it also needs to be updated in the UI code
	ErrNotActivated = status.Error(codes.Unimplemented, "the auth service is not activated")

	// ErrAlreadyActivated is returned by Activate if the Auth service
	// is already activated.
	ErrAlreadyActivated = status.Error(codes.Unimplemented, "the auth service is already activated")

	// ErrNotSignedIn indicates that the caller isn't signed in
	//
	// Note: This error message string is matched in the UI. If edited,
	// it also needs to be updated in the UI code
	ErrNotSignedIn = status.Error(codes.Unauthenticated, "no authentication token (try logging in)")

	// ErrNoMetadata is returned by the Auth API if the caller sent a request
	// containing no auth token.
	ErrNoMetadata = status.Error(codes.Internal, "no authentication metadata (try logging in)")

	// ErrBadToken is returned by the Auth API if the caller's token is corrupted
	// or has expired.
	ErrBadToken = status.Error(codes.Unauthenticated, "provided auth token is corrupted or has expired (try logging in again)")

	// ErrExpiredToken is returned by the Auth API if a restored token expired in
	// the past.
	ErrExpiredToken = status.Error(codes.Internal, "token expiration is in the past")
)

var DefaultOIDCScopes = []string{"email", "profile", "groups", oidc.ScopeOpenID}

// IsErrAlreadyActivated checks if an error is a ErrAlreadyActivated
func IsErrAlreadyActivated(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), status.Convert(ErrAlreadyActivated).Message())
}

// IsErrNotActivated checks if an error is a ErrNotActivated
func IsErrNotActivated(err error) bool {
	if err == nil {
		return false
	}
	// TODO(msteffen) This is unstructured because we have no way to propagate
	// structured errors across GRPC boundaries. Fix
	return strings.Contains(err.Error(), status.Convert(ErrNotActivated).Message())
}

// IsErrNotSignedIn returns true if 'err' is a ErrNotSignedIn
func IsErrNotSignedIn(err error) bool {
	if err == nil {
		return false
	}
	// TODO(msteffen) This is unstructured because we have no way to propagate
	// structured errors across GRPC boundaries. Fix
	return strings.Contains(err.Error(), status.Convert(ErrNotSignedIn).Message())
}

// IsErrNoMetadata returns true if 'err' is an ErrNoMetadata (uses string
// comparison to work across RPC boundaries)
func IsErrNoMetadata(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), status.Convert(ErrNoMetadata).Message())
}

// IsErrBadToken returns true if 'err' is a ErrBadToken
func IsErrBadToken(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), status.Convert(ErrBadToken).Message())
}

// IsErrExpiredToken returns true if 'err' is a ErrExpiredToken
func IsErrExpiredToken(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), status.Convert(ErrExpiredToken).Message())
}

const errNoRoleBindingMsg = "no role binding exists for"

// ErrNoRoleBinding is returned if no role binding exists for a resource.
type ErrNoRoleBinding struct {
	Resource Resource
}

func (e *ErrNoRoleBinding) Error() string {
	return fmt.Sprintf("%v %v %v", errNoRoleBindingMsg, e.Resource.Type, e.Resource.Name)
}

// IsErrNoRoleBinding checks if an error is a ErrNoRoleBinding
func IsErrNoRoleBinding(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), errNoRoleBindingMsg)
}

// ErrNotAuthorized is returned if the user is not authorized to perform
// a certain operation.
type ErrNotAuthorized struct {
	Subject string // subject trying to perform blocked operation -- always set

	Resource Resource     // Resource that the user is attempting to access
	Required []Permission // Caller needs 'Required'-level access to 'Resource'
}

// This error message string is matched in the UI. If edited,
// it also needs to be updated in the UI code
const errNotAuthorizedMsg = "not authorized to perform this operation"

func (e *ErrNotAuthorized) Error() string {
	return fmt.Sprintf("%v is %v - needs permissions %v on %v %v. Run `datactl auth roles-for-permission` to find roles that grant a given permission.", e.Subject, errNotAuthorizedMsg, e.Required, e.Resource.Type, e.Resource.Name)
}

// IsErrNotAuthorized checks if an error is a ErrNotAuthorized
func IsErrNotAuthorized(err error) bool {
	if err == nil {
		return false
	}
	// TODO(msteffen) This is unstructured because we have no way to propagate
	// structured errors across GRPC boundaries. Fix
	return strings.Contains(err.Error(), errNotAuthorizedMsg)
}

// ErrInvalidPrincipal indicates that a an argument to e.g. GetScope,
// SetScope, or SetACL is invalid
type ErrInvalidPrincipal struct {
	Principal string
}

func (e *ErrInvalidPrincipal) Error() string {
	return fmt.Sprintf("invalid principal \"%s\"; must start with one of \"pipeline:\", \"github:\", or \"robot:\", or have no \":\"", e.Principal)
}

// IsErrInvalidPrincipal returns true if 'err' is an ErrInvalidPrincipal
func IsErrInvalidPrincipal(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "invalid principal \"") &&
		strings.Contains(err.Error(), "\"; must start with one of \"pipeline:\", \"github:\", or \"robot:\", or have no \":\"")
}

// ErrTooShortTTL is returned by the ExtendAuthToken if request.Token already
// has a TTL longer than request.TTL.
type ErrTooShortTTL struct {
	RequestTTL, ExistingTTL int64
}

const errTooShortTTLMsg = "provided TTL (%d) is shorter than token's existing TTL (%d)"

func (e ErrTooShortTTL) Error() string {
	return fmt.Sprintf(errTooShortTTLMsg, e.RequestTTL, e.ExistingTTL)
}

// IsErrTooShortTTL returns true if 'err' is a ErrTooShortTTL
func IsErrTooShortTTL(err error) bool {
	if err == nil {
		return false
	}
	errMsg := err.Error()
	return strings.Contains(errMsg, "provided TTL (") &&
		strings.Contains(errMsg, ") is shorter than token's existing TTL (") &&
		strings.Contains(errMsg, ")")
}

// HashToken converts a token to a cryptographic hash.
// We don't want to store tokens verbatim in the database, as then whoever
// that has access to the database has access to all tokens.
func HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return fmt.Sprintf("%x", sum)
}

// GetAuthToken extracts the auth token embedded in 'ctx', if there is one
func GetAuthToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.EnsureStack(ErrNoMetadata)
	}
	if len(md[ContextTokenKey]) > 1 {
		return "", errors.Errorf("multiple authentication token keys found in context")
	} else if len(md[ContextTokenKey]) == 0 {
		return "", errors.EnsureStack(ErrNotSignedIn)
	}
	return md[ContextTokenKey][0], nil
}
