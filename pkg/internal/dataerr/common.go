package dataerr

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
	"fmt"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/bhojpur/data/pkg/internal/errors"
)

type ErrNotExist struct {
	Collection string
	ID         string
}

func NewNotExist(collection, id string) error {
	return ErrNotExist{
		Collection: collection,
		ID:         id,
	}
}

func (e ErrNotExist) Error() string {
	return fmt.Sprintf("%s does not contain item: (%s)", e.Collection, e.ID)
}

func (e ErrNotExist) GRPCStatus() *status.Status {
	return status.New(codes.NotFound, e.Error())
}

func IsNotExist(err error) bool {
	target := ErrNotExist{}
	return errors.As(err, &target) || os.IsNotExist(err) || os.IsNotExist(errors.Unwrap(err))
}

type ErrExists struct {
	Collection string
	ID         string
}

func NewExists(collection, id string) error {
	return &ErrExists{
		Collection: collection,
		ID:         id,
	}
}

func (e ErrExists) Error() string {
	return fmt.Sprintf("%s already contains an item: (%s)", e.Collection, e.ID)
}

func (e ErrExists) GRPCStatus() *status.Status {
	return status.New(codes.AlreadyExists, e.Error())
}

func IsExists(err error) bool {
	target := ErrExists{}
	return errors.As(err, &target)
}

var (
	// ErrBreak is an error used to break out of call back based iteration,
	// should be swallowed by iteration functions and treated as successful
	// iteration.
	ErrBreak = errors.Errorf("BREAK")
)
