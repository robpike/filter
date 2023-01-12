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

func fac(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fac(n-1)
}

func TestReduce(t *testing.T) {
	const size = 10
	a := make([]int, size)
	for i := range a {
		a[i] = i + 1
	}
	for i := 1; i < 10; i++ {
		out := Reduce(a[:i], mul, 1)
		expect := fac(i)
		if expect != out {
			t.Fatalf("expected %d got %d", expect, out)
		}
	}
}
