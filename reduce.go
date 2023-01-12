// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filter

// Reduce computes the reduction of the pair function across the elements of
// the slice. For instance, if the slice contains successive integers starting
// at 1 and the function is multiply, the result will be the factorial function.
// If the slice is empty, Reduce returns zero; if it has only one element, it
// returns that element. Example:
//
//	func multiply(a, b int) int { return a*b }
//	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	factorial := Reduce(a, multiply, 1).(int)
func Reduce[T any](slice []T, pairFunction func(T, T) T, zero T) T {
	n := len(slice)
	if n == 0 {
		return zero
	}
	if n == 1 {
		return slice[0]
	}
	out := slice[0] // By convention, pairFunction(zero, slice[0]) = slice[0].
	// Run from index 1 to the end.
	for i := 1; i < n; i++ {
		out = pairFunction(out, slice[i])
	}
	return out
}
