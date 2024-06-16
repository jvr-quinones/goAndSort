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

func radixSortMSD(slice []int, base int) (sorted []int) {
	if len(slice) < 2 {
		return slice
	}

	digits := make([][]int, base)
	sorted = make([]int, 0)
	maxLogN := 0
	ind1 := 0

	for _, val := range slice {
		logN := 0.0
		if base == 2 {
			logN = math.Log2(float64(val))
		} else if base == 10 {
			logN = math.Log10(float64(val))
		} else {
			logN = math.Log10(float64(val)) / math.Log10(float64(base))
		}

		if maxLogN < int(logN) {
			maxLogN = int(logN)
		}
	}

	for ; maxLogN >= 0; maxLogN-- {
		tryAgain := false
		for ind1 = range slice {
			divisor := int(math.Pow(float64(base), float64(maxLogN)))
			remain := (slice[ind1] / divisor) % base
			digits[remain] = append(digits[remain], slice[ind1])
		}

		for ind1 = range digits {
			tryAgain = tryAgain || len(digits[ind1]) == len(slice)
			if !tryAgain {
				digits[ind1] = radixSortMSD(digits[ind1], base)
			} else {
				digits[ind1] = make([]int, 0)
			}
			sorted = slices.Concat(sorted, digits[ind1])
		}

		if len(sorted) > len(slice) {
			logger.Fatal("Mistakes were made: the sorted array is bigger than the original :O")
		}

		if !tryAgain {
			break
		}
	}

	return sorted
}
