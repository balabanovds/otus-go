package max

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

	return nil
}