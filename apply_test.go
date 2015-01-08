// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filter

import (
	"reflect"
	"testing"
)

func triple(a int) int {
	return a * 3
}

func tripleString(a string) string {
	return a + a + a
}

func tripleToFloat(a int) float64 {
	return float64(a * 3)
}

func floatToInt(a float64) int {
	return int(a)
}

func isEven(a int) bool {
	return a%2 == 0
}

func isEvenString(a string) bool {
	return a[0]%2 == 0
}

func is18(a int) bool {
	return a == 18
}

func TestApply(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := []int{3, 6, 9, 12, 15, 18, 21, 24, 27}
	result := Apply(a, triple)
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("Apply failed: expect %v got %v", expect, result)
	}
}

func TestApplyString(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	expect := []string{"111", "222", "333", "444", "555", "666", "777", "888", "999"}
	result := Apply(a, tripleString)
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("Apply failed: expect %v got %v", expect, result)
	}
}

func TestApplyDifferentTypes(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := []int{3, 6, 9, 12, 15, 18, 21, 24, 27}
	result := Apply(Apply(a, tripleToFloat), floatToInt)
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("Apply failed: expect %v got %v", expect, result)
	}
}

func TestChoose(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := []int{2, 4, 6, 8}
	result := Choose(a, isEven)
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("Choose failed: expect %v got %v", expect, result)
	}
}

func TestDrop(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := []int{1, 3, 5, 7, 9}
	result := Drop(a, isEven)
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("Drop failed: expect %v got %v", expect, result)
	}
}

func TestApplyInPlace(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := []int{3, 6, 9, 12, 15, 18, 21, 24, 27}
	ApplyInPlace(a, triple)
	if !reflect.DeepEqual(expect, a) {
		t.Fatalf("Apply failed: expect %v got %v", expect, a)
	}
}

func TestApplyInPlaceString(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	expect := []string{"111", "222", "333", "444", "555", "666", "777", "888", "999"}
	ApplyInPlace(a, tripleString)
	if !reflect.DeepEqual(expect, a) {
		t.Fatalf("Apply failed: expect %v got %v", expect, a)
	}
}

func TestChooseInPlaceString(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	expect := []string{"2", "4", "6", "8"}
	ChooseInPlace(&a, isEvenString)
	if !reflect.DeepEqual(expect, a) {
		t.Fatalf("ChooseInPlace failed: expect %v got %v", expect, a)
	}
}

func TestDropInPlaceString(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	expect := []string{"1", "3", "5", "7", "9"}
	DropInPlace(&a, isEvenString)
	if !reflect.DeepEqual(expect, a) {
		t.Fatalf("DropInPlace failed: expect %v got %v", expect, a)
	}
}
