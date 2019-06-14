package hw04

// FindMax finds maximum in given slice based on incoming predicate
func FindMax(incoming interface{}, predicate func(i, j interface{}) bool) interface{} {
	slice := make([]interface{}, len(incoming))

	for i, v := range incoming {
		slice[i] = v
	}

	max := slice[0]

	for _, v := range slice {
		if predicate(max, v) {
			max = v
		}
	}

	return max
}
