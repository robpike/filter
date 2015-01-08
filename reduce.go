// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filter

import (
	"reflect"
)

// Reduce computes the reduction of the pair function across the elements of
// the slice. (If the types of the slice and function do not correspond, Reduce
// panics.) For instance, if the slice contains successive integers starting at
// 1 and the function is multiply, the result will be the factorial function.
// If the slice is empty, Reduce returns zero; if it has only one element, it
// returns that element. The return value must be type-asserted by the caller
// back to the element type of the slice. Example:
//	func multiply(a, b int) int { return a*b }
//	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	factorial := Reduce(a, multiply, 1).(int)
func Reduce(slice, pairFunction, zero interface{}) interface{} {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("reduce: not slice")
	}
	n := in.Len()
	switch n {
	case 0:
		return zero
	case 1:
		return in.Index(0)
	}
	elemType := in.Type().Elem()
	fn := reflect.ValueOf(pairFunction)
	if !goodFunc(fn, elemType, elemType, elemType) {
		str := elemType.String()
		panic("apply: function must be of type func(" + str + ", " + str + ") " + str)
	}
	// Do the first two by hand to prime the pump.
	var ins [2]reflect.Value
	ins[0] = in.Index(0)
	ins[1] = in.Index(1)
	out := fn.Call(ins[:])[0]
	// Run from index 2 to the end.
	for i := 2; i < n; i++ {
		ins[0] = out
		ins[1] = in.Index(i)
		out = fn.Call(ins[:])[0]
	}
	return out.Interface()
}
