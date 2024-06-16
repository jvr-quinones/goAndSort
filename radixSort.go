package main

import (
	"math"
	"slices"
)

func radixSortLSD(slice []int, base int) (sorted []int) {
	if base < 2 {
		logger.Fatal("Base has to be greater than 2")
	}

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

		for ind2 = range digits {
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
	if base < 2 {
		logger.Fatal("Base has to be greater than 2")
	} else if len(slice) < 2 {
		return slice
	}

	exp := 0
	ind1 := 1
	sliceMin := slice[0]
	sliceMax := slice[0]

	for _, val := range slice {
		if val < sliceMin {
			sliceMin = val
		} else if val > sliceMax {
			sliceMax = val
		}
	}

	sliceDiff := sliceMax - sliceMin
	if sliceDiff == 0 {
		return slice
	} else if sliceDiff < base {
		exp = sliceMax/base - sliceMin/base
	} else if base == 2 {
		exp = int(math.Log2(float64(sliceDiff)))
	} else if base == 10 {
		exp = int(math.Log10(float64(sliceDiff)))
	} else {
		exp = int(math.Log10(float64(sliceDiff)) / math.Log10(float64(base)))
	}

	digits := make([][]int, base)
	sorted = make([]int, 0)
	for ind1 = range slice {
		divisor := int(math.Pow(float64(base), float64(exp)))
		remain := (slice[ind1] / divisor) % base
		digits[remain] = append(digits[remain], slice[ind1])
	}

	for ind1 = range digits {
		digits[ind1] = radixSortMSD(digits[ind1], base)
		sorted = slices.Concat(sorted, digits[ind1])
	}

	if len(sorted) > len(slice) {
		logger.Fatal("Mistakes were made: the sorted array is bigger than the original :O")
	}

	return sorted
}
