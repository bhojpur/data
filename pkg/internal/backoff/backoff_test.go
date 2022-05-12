package backoff_test

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
	"time"

	"github.com/bhojpur/data/pkg/internal/backoff"
	"github.com/bhojpur/data/pkg/internal/errors"
)

func TestNextBackOffMillis(t *testing.T) {
	subtestNextBackOff(t, 0, new(backoff.ZeroBackOff))
	subtestNextBackOff(t, backoff.Stop, new(backoff.StopBackOff))
}

func subtestNextBackOff(t *testing.T, expectedValue time.Duration, backOffPolicy backoff.BackOff) {
	for i := 0; i < 10; i++ {
		next := backOffPolicy.NextBackOff()
		if next != expectedValue {
			t.Errorf("got: %d expected: %d", next, expectedValue)
		}
	}
}

func TestConstantBackOff(t *testing.T) {
	backoff := backoff.NewConstantBackOff(time.Second)
	if backoff.NextBackOff() != time.Second {
		t.Error("invalid interval")
	}
}

func abstime(t time.Duration) time.Duration {
	if t < 0 {
		return -t
	}
	return t
}

func TestConstantBackOffCompare(t *testing.T) {
	var callTimes [10]time.Time
	idx := 0
	start := time.Now()
	err := backoff.Retry(func() error {
		callTimes[idx] = time.Now()
		idx++
		return errors.Errorf("expected error")
	}, backoff.RetryEvery(time.Second).For(9*time.Second))
	if err.Error() != "expected error" {
		t.Fatalf("Retry loop didn't return internal error to caller")
	}

	epsilon := 500 * time.Millisecond
	if idx < 8 {
		t.Fatalf("expected 9 retries, but only saw %d", idx)
	}
	nextT := start
	for i := 0; i < idx; i++ {
		if abstime(callTimes[i].Sub(nextT)) > epsilon {
			t.Fatalf("expected retry %d to occur at %s but actually occurred at %s",
				i, nextT.String(), callTimes[i].String())
		}
		nextT = nextT.Add(1 * time.Second)
	}
}
