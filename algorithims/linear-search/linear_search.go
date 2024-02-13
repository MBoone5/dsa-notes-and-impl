package search

func linearSearch(haystack []float64, needle float64) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return i
		}
	}

	return -1
}
