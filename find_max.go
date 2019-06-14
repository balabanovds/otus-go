package hw04

// FindMax finds maximum in given slice based on incoming predicate
func FindMax(incoming []interface{}, predicate func(i, j interface{}) bool) interface{} {

	max := incoming[0]

	for _, v := range incoming {
		if predicate(max, v) {
			max = v
		}
	}

	return max
}
