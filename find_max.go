package max

import (
	"reflect"
)

// FindMax finds max of values incoming in slice structure,
// based on comparator provided by user
func FindMax(slice []interface{}, less func(i, j int) bool) interface{} {
	s := make([]interface{}, len(slice))
	for i, v := range slice {
		s[i] = v
	}

	max := s[0]
	for i := 1; i < len(s); i++ {
		if less(i-1, i) {
			max = s[i]
		}
	}

	return max
}

// FindMaxReflection finds max of values in slice structure,
// based on comparator provided by user, but slice is provided as interface{}
func FindMaxReflection(slice interface{}, less func(i, j int) bool) interface{} {

	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		panic("Income is not a slice")
	}
	s := reflect.ValueOf(slice)

	max := conv(s.Index(0))
	for i := 1; i < s.Len(); i++ {
		if less(i-1, i) {
			max = conv(s.Index(i))
		}
	}

	return max
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
