package main

import (
	"math"
	"slices"
)

func radixSortLSD(slice []int, base int) (sorted []int) {
	digits := make([][]int, base)
	maxLogN := 0
	ind1 := 0

	for _, val := range slice {
		logN := 0
		if base == 2 {
			logN = int(math.Log2(float64(val)))
		} else if base == 10 {
			logN = int(math.Log10(float64(val)))
		} else {
			logN = int(math.Log10(float64(val)) / math.Log10(float64(base)))
		}

		if maxLogN < int(logN) {
			maxLogN = int(logN)
		}
	}

	for ind1 = 0; ind1 <= maxLogN; ind1++ {
		ind2 := 0
		if ind1 == 0 {
			sorted = make([]int, len(slice))
			copy(sorted, slice)
		}

		for ind2 = range sorted {
			divisor := int(math.Pow(float64(base), float64(ind1)))
			remain := (sorted[ind2] / divisor) % base
			digits[remain] = append(digits[remain], sorted[ind2])
		}

		for ind2 = 0; ind2 < base; ind2++ {
			if ind2 == 0 {
				sorted = make([]int, 0)
			}
			sorted = slices.Concat(sorted, digits[ind2])
			digits[ind2] = make([]int, 0)
		}
	}

	return sorted
}
