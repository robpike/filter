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

func TestReduceMul(t *testing.T) {
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

func TestReduceMul2(t *testing.T) {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}
	// Compute 10!
	out := Reduce(a, mul, 2).(int)
	expect := 2
	for i := range a {
		expect *= a[i]
	}
	if expect != out {
		t.Fatalf("expected %d got %d", expect, out)
	}
}

func k(x1, x2 *int) *int {
	return x1
}

func TestReduceNil(t *testing.T) {
	a := make([]*int, 10)
	for i := range a {
		a[i] = nil
	}
	var zero *int = nil
	// Compute nil!
	out := Reduce(a, k, zero).(*int)
	expect := zero
	if expect != out {
		t.Fatalf("expected %p got %p", expect, out)
	}
}

type IntElement struct {
	Value int
	Next *IntElement
}

type IntList *IntElement

func IntListEquals(l1, l2 IntList) bool {
	for ;l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		if l1.Value != l2.Value {
			return false
		}
	}
	return true
}

func rCons(is IntList, i int) IntList {
	return &IntElement{i, is}
}

var list4321 = rCons(rCons(rCons(rCons(nil, 1), 2), 3), 4)

func TestReduceIntList(t *testing.T) {
	a := make([]int, 4)
	for i := range a {
		a[i] = i + 1
	}
	var zero IntList = nil
	out := Reduce(a, rCons, zero).(IntList)
	if !IntListEquals(out, list4321) {
		t.Fatalf("expected [4,3,2,1] got something else")
	}
}
