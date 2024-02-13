package search

import (
	"math"
)

func binarySearch(haystack []float64, needle float64) int {
	// Assuming the haystack is sorted

	var low float64 = 0
	var high float64 = float64(len(haystack))

	for low < high {
    var m int
    
		m = int(math.Floor((low + (high - low)) / 2))

		if haystack[m] == needle {
			return m
		}

		if haystack[m] > needle {
			high = float64(m)
		}

		if haystack[m] < needle {
			low = float64(m) + 1
		}
	}

	return -1
}
