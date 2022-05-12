package errors

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
	"reflect"

	"github.com/pkg/errors"
)

var errorType = reflect.TypeOf((*error)(nil)).Elem()

func tryAs(err error, targetVal reflect.Value) bool {
	e := targetVal.Type().Elem()
	if e.Kind() == reflect.Interface || e.Implements(errorType) {
		return errors.As(err, targetVal.Interface())
	}
	return false
}

// As finds the first error in err's chain that matches the target's type, and
// if so, sets target to that error value and returns true.
// As is a wrapper for the underlying errors.As function, which may panic or
// return unexpected results based on how err was constructed (with or without
// a pointer).  This works by inspecting the type of target and attempting
// multiple errors.As calls if necessary.
func As(err error, target interface{}) bool {
	v := reflect.ValueOf(target)

	switch v.Kind() {
	case reflect.Ptr:
		// Attempt unwrapping a nested pointer
		if v.Elem().Kind() == reflect.Ptr && tryAs(err, v.Elem()) {
			return true
		}

		// Attempt wrapping with an additional pointer
		vp := reflect.New(v.Type())
		if tryAs(err, vp) {
			v.Elem().Set(vp.Elem().Elem())
			return true
		}

		// Attempt the passed target as-is
		return tryAs(err, v)
	}
	panic(fmt.Sprintf("unexpected target type: %v, errors.As must be passed a pointer target", v.Kind()))
}
