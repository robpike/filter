// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filter

import (
	"testing"
)

func mul(a, b int) int {
	return a * b
}

func TestReduce(t *testing.T) {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}
	// Compute 10!
	out := Reduce(a, mul, 1).(int)
	expect := 1
	for i := range a {
		expect *= a[i]
	}
	if expect != out {
		t.Fatalf("expected %d got %d", expect, out)
	}
}
