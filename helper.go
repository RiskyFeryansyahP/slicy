package main

import "reflect"

// IsFunction checks whether the type used is a function and has 1 argument and 1 return
func IsFunction(fn interface{}) bool {
	fnType := reflect.TypeOf(fn)

	res := fnType.Kind() == reflect.Func && fnType.NumIn() == 1 && fnType.NumOut() == 1

	return res
}
