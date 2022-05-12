package errutil

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
	"net"
	"strings"

	"github.com/bhojpur/data/pkg/internal/dataerr"
	"github.com/bhojpur/data/pkg/internal/errors"
)

var ErrBreak = dataerr.ErrBreak

// IsAlreadyExistError returns true if err is due to trying to create a
// resource that already exists. It uses simple string matching, it's not
// terribly smart.
func IsAlreadyExistError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "already exists")
}

// IsNotFoundError returns true if err is due to a resource not being found. It
// uses simple string matching, it's not terribly smart.
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "not found")
}

// IsWriteToOutputBranchError returns true if the err is due to an attempt to
// write to an output repo/branch
func IsWriteToOutputBranchError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "cannot start a commit on an output branch")
}

// IsNotADirectoryError returns true if the err is due to an attempt to put a
// file on path that has a non-directory parent. These errors come from the
// hashtree package; while it provides an error code, we can't check against
// that because we'd then have to import hashtree, and hashtree imports
// errutil, leading to circular imports.
func IsNotADirectoryError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "but it's not a directory")
}

// IsInvalidPathError returns true if the err is due to an invalid path
func IsInvalidPathError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "only printable ASCII characters allowed") ||
		strings.Contains(err.Error(), "not allowed in path")
}

// IsNetRetryable returns true if the error is a temporary network error.
func IsNetRetryable(err error) bool {
	var netErr net.Error
	return errors.As(err, &netErr) && netErr.Temporary()
}
