package datahash

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
	"encoding/hex"
	"hash"

	"github.com/bhojpur/data/pkg/internal/errors"
	"golang.org/x/crypto/blake2b"
)

// OutputSize is the size of an Output in bytes
const OutputSize = 32

// Output is the output of the hash function.
// Sum returns an Output
type Output = [OutputSize]byte

// New creates a new hasher.
func New() hash.Hash {
	h, err := blake2b.New256(nil)
	if err != nil {
		panic(err)
	}
	return h
}

// Sum computes a hash sum for a set of bytes.
func Sum(data []byte) Output {
	return blake2b.Sum256(data)
}

// ParseHex parses a hex string into output.
func ParseHex(x []byte) (*Output, error) {
	o := Output{}
	n, err := hex.Decode(o[:], x)
	if err != nil {
		return nil, errors.EnsureStack(err)
	}
	if n < OutputSize {
		return nil, errors.Errorf("hex string too short to be Output")
	}
	return &o, nil
}

// EncodeHash encodes a hash into a string representation.
func EncodeHash(bytes []byte) string {
	return hex.EncodeToString(bytes)
}
