// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package filter contains utility functions for filtering slices through the
// distributed application of a filter function.
//
// The package is an experiment to see how easy it is to write such things
// in Go. It is easy, but for loops are just as easy and more efficient.
// 
// You should not use this package.
//
package filter // import "robpike.io/filter"

import (
	"reflect"
)

// Apply takes a slice of type []T and a function of type func(T) T. It
// returns a newly allocated slice where each element is the result of calling
// the function on successive elements of the slice.
func Apply[T any, U any](slice []T, function func(T) U) []U {
	out := make([]U, len(slice), len(slice))
	for i := 0; i < len(slice); i++ {
		out[i] = function(slice[i])
	}
	return out
}

// ApplyInPlace is like Apply, but overwrites the slice rather than returning a
// newly allocated slice.
func ApplyInPlace[T any](slice []T, function func(T) T) {
	for i := 0; i < len(slice); i++ {
		slice[i] = function(slice[i])
	}
}

// Choose takes a slice of type []T and a function of type func(T) bool. It
// returns a newly allocated slice containing only those elements of the input
// slice that satisfy the function.
func Choose[T any](slice []T, function func(T) bool) []T {
	return chooseOrDrop(slice, function, true)
}

// Drop takes a slice of type []T and a function of type func(T) bool. It
// returns a newly allocated slice containing only those elements of the input
// slice that do not satisfy the function, that is, it removes elements that
// satisfy the function.
func Drop[T any](slice []T, function func(T) bool) []T {
	return chooseOrDrop(slice, function, false)
}

func chooseOrDrop[T any](slice []T, function func(T) bool, truth bool) []T {
	var out []T
	for i := 0; i < len(slice); i++ {
		if truth == function(slice[i]) {
			out = append(out, slice[i])
		}
	}
	return out
}

// ChooseInPlace is like Choose, but overwrites the slice rather than returning
// a newly allocated slice. Since ChooseInPlace must modify the header of the
// slice to set the new length, it takes as argument a pointer to a slice
// rather than a slice.
func ChooseInPlace[T any](pointerToSlice *[]T, function func(T) bool) {
	chooseOrDropInPlace(pointerToSlice, function, true)
}

// DropInPlace is like Drop, but overwrites the slice rather than returning a
// newly allocated slice. Since DropInPlace must modify the header of the slice
// to set the new length, it takes as argument a pointer to a slice rather than
// a slice.
func DropInPlace[T any](pointerToSlice *[]T, function func(T) bool) {
	chooseOrDropInPlace(pointerToSlice, function, false)
}

func chooseOrDropInPlace[T any](pointerToSlice *[]T, function func(T) bool, truth bool) {
	n := 0
	for i := 0; i < len(*pointerToSlice); i++ {
		if truth == function((*pointerToSlice)[i]) {
			(*pointerToSlice)[n] = (*pointerToSlice)[i]
			n++
		}
	}
	reflect.ValueOf(pointerToSlice).Elem().SetLen(n)
}
