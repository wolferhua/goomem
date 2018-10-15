package goomem

import (
	"reflect"
)

func isBaseType(val interface{}) bool {
	kind := reflect.TypeOf(val).Kind()
	return kind <= reflect.Float64 || kind == reflect.String
}
