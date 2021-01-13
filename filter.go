package main

import (
	"reflect"
)

// Filter iterates over elements of collection, returning an array of
func Filter(slice, fn interface{}) interface{} {
	sliceValue := reflect.ValueOf(slice)
	sliceType := reflect.TypeOf(slice)

	fnValue := reflect.ValueOf(fn)

	// get the slice type of the array type passed in the first argument
	resultSliceType := reflect.SliceOf(sliceType.Elem())

	// create a new slice to be the return value
	resultSlice := reflect.MakeSlice(resultSliceType, 0, 0)

	if sliceType.Kind() != reflect.Slice {
		panic("first argument must be slice")
	}

	if !IsFunction(fn) {
		panic("first argument must be function with 1 argument and return type bool")
	}

	for i := 0; i < sliceValue.Len(); i++ {
		elem := sliceValue.Index(i)

		result := fnValue.Call([]reflect.Value{elem})[0].Interface().(bool)

		if result {
			resultSlice = reflect.Append(resultSlice, elem)
		}
	}

	return resultSlice
}
