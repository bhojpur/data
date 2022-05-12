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
	"math"
	"testing"
	"time"

	"github.com/bhojpur/data/pkg/internal/backoff"
)

func TestBackOff(t *testing.T) {
	var (
		testInitialInterval     = 500 * time.Millisecond
		testRandomizationFactor = 0.1
		testMultiplier          = 2.0
		testMaxInterval         = 5 * time.Second
		testMaxElapsedTime      = 15 * time.Minute
	)

	exp := backoff.NewExponentialBackOff()
	exp.InitialInterval = testInitialInterval
	exp.RandomizationFactor = testRandomizationFactor
	exp.Multiplier = testMultiplier
	exp.MaxInterval = testMaxInterval
	exp.MaxElapsedTime = testMaxElapsedTime
	exp.Reset()

	var expectedResults = []time.Duration{500, 1000, 2000, 4000, 5000, 5000, 5000, 5000, 5000, 5000}
	for i, d := range expectedResults {
		expectedResults[i] = d * time.Millisecond
	}

	for _, expected := range expectedResults {
		assertEquals(t, expected, exp.CurrentInterval)
		// Assert that the next backoff falls in the expected range.
		var minInterval = expected - time.Duration(testRandomizationFactor*float64(expected))
		var maxInterval = expected + time.Duration(testRandomizationFactor*float64(expected))
		var actualInterval = exp.NextBackOff()
		if !(minInterval <= actualInterval && actualInterval <= maxInterval) {
			t.Error("error")
		}
	}
}

func TestGetRandomizedInterval(t *testing.T) {
	// 33% chance of being 1.
	assertEquals(t, 1, backoff.GetRandomValueFromInterval(0.5, 0, 2))
	assertEquals(t, 1, backoff.GetRandomValueFromInterval(0.5, 0.33, 2))
	// 33% chance of being 2.
	assertEquals(t, 2, backoff.GetRandomValueFromInterval(0.5, 0.34, 2))
	assertEquals(t, 2, backoff.GetRandomValueFromInterval(0.5, 0.66, 2))
	// 33% chance of being 3.
	assertEquals(t, 3, backoff.GetRandomValueFromInterval(0.5, 0.67, 2))
	assertEquals(t, 3, backoff.GetRandomValueFromInterval(0.5, 0.99, 2))
}

type TestClock struct {
	i     time.Duration
	start time.Time
}

func (c *TestClock) Now() time.Time {
	t := c.start.Add(c.i)
	c.i += time.Second
	return t
}

func TestGetElapsedTime(t *testing.T) {
	var exp = backoff.NewExponentialBackOff()
	exp.Clock = &TestClock{}
	exp.Reset()

	var elapsedTime = exp.GetElapsedTime()
	if elapsedTime != time.Second {
		t.Errorf("elapsedTime=%d", elapsedTime)
	}
}

func TestMaxElapsedTime(t *testing.T) {
	var exp = backoff.NewExponentialBackOff()
	exp.Clock = &TestClock{start: time.Time{}.Add(10000 * time.Second)}
	// Change the currentElapsedTime to be 0 ensuring that the elapsed time will be greater
	// than the max elapsed time.
	exp.StartTime = time.Time{}
	assertEquals(t, backoff.Stop, exp.NextBackOff())
}

func TestBackOffOverflow(t *testing.T) {
	var (
		testInitialInterval time.Duration = math.MaxInt64 / 2
		testMaxInterval     time.Duration = math.MaxInt64
		testMultiplier                    = 2.1
	)

	exp := backoff.NewExponentialBackOff()
	exp.InitialInterval = testInitialInterval
	exp.Multiplier = testMultiplier
	exp.MaxInterval = testMaxInterval
	exp.Reset()

	exp.NextBackOff()
	// Assert that when an overflow is possible the current varerval   time.Duration    is set to the max varerval   time.Duration   .
	assertEquals(t, testMaxInterval, exp.CurrentInterval)
}

func assertEquals(t *testing.T, expected, value time.Duration) {
	if expected != value {
		t.Errorf("got: %d, expected: %d", value, expected)
	}
}
