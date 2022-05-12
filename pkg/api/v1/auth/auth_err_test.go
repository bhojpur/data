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
	"testing"

	"github.com/bhojpur/data/pkg/internal/errors"
	"github.com/bhojpur/data/pkg/internal/require"
)

// grpcify returns an error e such that e.Error() is similar to what grpc
// errors emit (though this doesn't convert 'error' to an actual GRPC error)
func grpcify(err error) error {
	return errors.Errorf("rpc error: code = Unknown desc = %s", err.Error())
}

func TestIsErrNotActivated(t *testing.T) {
	require.False(t, IsErrNotActivated(nil))
	require.True(t, IsErrNotActivated(ErrNotActivated))
	require.True(t, IsErrNotActivated(grpcify(ErrNotActivated)))
}

func TestIsErrAlreadyActivated(t *testing.T) {
	require.False(t, IsErrAlreadyActivated(nil))
	require.True(t, IsErrAlreadyActivated(ErrAlreadyActivated))
	require.True(t, IsErrAlreadyActivated(grpcify(ErrAlreadyActivated)))
}

func TestIsErrNotSignedIn(t *testing.T) {
	require.False(t, IsErrNotSignedIn(nil))
	require.True(t, IsErrNotSignedIn(ErrNotSignedIn))
	require.True(t, IsErrNotSignedIn(grpcify(ErrNotSignedIn)))
}

func TestIsErrNoMetadata(t *testing.T) {
	require.False(t, IsErrNoMetadata(nil))
	require.True(t, IsErrNoMetadata(ErrNoMetadata))
	require.True(t, IsErrNoMetadata(grpcify(ErrNoMetadata)))
}

func TestIsErrBadToken(t *testing.T) {
	require.False(t, IsErrBadToken(nil))
	require.True(t, IsErrBadToken(ErrBadToken))
	require.True(t, IsErrBadToken(grpcify(ErrBadToken)))
}

func TestIsErrNotAuthorized(t *testing.T) {
	require.False(t, IsErrNotAuthorized(nil))
	require.True(t, IsErrNotAuthorized(&ErrNotAuthorized{
		Subject:  "alice",
		Resource: Resource{Type: ResourceType_REPO, Name: "data"},
		Required: []Permission{},
	}))
	require.True(t, IsErrNotAuthorized(grpcify(&ErrNotAuthorized{
		Subject:  "alice",
		Resource: Resource{Type: ResourceType_REPO, Name: "data"},
		Required: []Permission{},
	})))
	require.True(t, IsErrNotAuthorized(&ErrNotAuthorized{
		Subject:  "alice",
		Resource: Resource{Type: ResourceType_CLUSTER},
		Required: []Permission{},
	}))
	require.True(t, IsErrNotAuthorized(grpcify(&ErrNotAuthorized{
		Subject:  "alice",
		Resource: Resource{Type: ResourceType_CLUSTER},
		Required: []Permission{},
	})))
}

func TestErrNoRoleBinding(t *testing.T) {
	require.True(t, IsErrNoRoleBinding(&ErrNoRoleBinding{
		Resource{Type: ResourceType_REPO, Name: "test"},
	}))
	require.True(t, IsErrNoRoleBinding(grpcify(&ErrNoRoleBinding{
		Resource{Type: ResourceType_REPO, Name: "test"},
	})))
}

func TestIsErrInvalidPrincipal(t *testing.T) {
	require.False(t, IsErrInvalidPrincipal(nil))
	require.True(t, IsErrInvalidPrincipal(&ErrInvalidPrincipal{
		Principal: "alice",
	}))
	require.True(t, IsErrInvalidPrincipal(grpcify(&ErrInvalidPrincipal{
		Principal: "alice",
	})))
}

func TestIsErrTooShortTTL(t *testing.T) {
	require.False(t, IsErrTooShortTTL(nil))
	require.True(t, IsErrTooShortTTL(ErrTooShortTTL{
		RequestTTL:  1234,
		ExistingTTL: 2345,
	}))
	require.True(t, IsErrTooShortTTL(grpcify(ErrTooShortTTL{
		RequestTTL:  1234,
		ExistingTTL: 2345,
	})))
}
