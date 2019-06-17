package max

import (
	"fmt"
	"reflect"
)

// FindMax finds max of values incoming in slice structure,
// based on comparator provided by user
func FindMax(slice []interface{}, less func(i, j int) bool) interface{} {
	if len(slice) == 0 {
		return nil
	}

	max := slice[0]
	l := len(slice)

	for i := 1; i < l; i++ {
		if less(i-1, i) {
			max = slice[i]
		}
	}

	return max
}

// FindMaxReflection finds max of values in slice structure,
// based on comparator provided by user, but slice is provided as interface{}
func FindMaxReflection(slice interface{}, less func(i, j int) bool) (interface{}, error) {

	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return nil, fmt.Errorf("First argument should be slice")
	}
	s := reflect.ValueOf(slice)

	max := conv(s.Index(0))
	for i := 1; i < s.Len(); i++ {
		if less(i-1, i) {
			max = conv(s.Index(i))
		}
	}

	return max, nil
}

func conv(v reflect.Value) interface{} {
	switch v.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.String:
		return v.String()
	case reflect.Bool:
		return v.Bool()
	case reflect.Float32, reflect.Float64:
		return v.Float()
	}
	return nil
}
