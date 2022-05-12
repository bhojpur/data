package identity

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
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrInvalidID is returned if the client or connector ID does not exist
	ErrInvalidID = status.Error(codes.Internal, "ID does not exist")

	// ErrAlreadyExists is returned if the client or connector ID already exists
	ErrAlreadyExists = status.Error(codes.Internal, "ID already exists")
)

// IsErrInvalidID checks if an error is a ErrInvalidID
func IsErrInvalidID(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), status.Convert(ErrInvalidID).Message())
}

// IsErrAlreadyExists checks if an error is a ErrAlreadyExists
func IsErrAlreadyExists(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), status.Convert(ErrAlreadyExists).Message())
}
