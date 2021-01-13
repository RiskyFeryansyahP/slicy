package main

import (
	"reflect"
	"strings"
)

// Find iterates over elements of collection, returning the first
func Find(i, fn interface{}) interface{} {
	typeOfI := reflect.TypeOf(i).Kind()
	valueOfI := reflect.ValueOf(i)

	funcValue := reflect.ValueOf(fn)

	if typeOfI != reflect.Slice {
		panic("first arg must be slice")
	}

	if !IsFunction(fn) {
		panic("first argument must be function with 1 argument and return type bool")
	}

	for i := 0; i < valueOfI.Len(); i++ {
		elem := valueOfI.Index(i)
		result := funcValue.Call([]reflect.Value{elem})[0].Interface().(bool)

		if result {
			return valueOfI.Index(i).Interface()
		}
	}

	return nil
}

// IndexOf gets the index at which the first occurrence of value is found in array or return -1
func IndexOf(slice, elemen interface{}) int {
	sliceValue := reflect.ValueOf(slice)
	elemenValue := reflect.ValueOf(elemen)

	sliceType := reflect.TypeOf(slice)

	if sliceType.Kind() == reflect.String {
		strings.Index(sliceValue.String(), elemenValue.String())
	}

	if sliceType.Kind() == reflect.Slice {
		for i := 0; i < sliceValue.Len(); i++ {
			if sliceValue.Index(i).Interface() == elemen {
				return i
			}
		}
	}

	return -1
}
