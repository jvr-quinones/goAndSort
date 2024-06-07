package main

import (
	"slices"
)

func quickSortInPlace(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	// Partition
	ind1 := 1
	ind2 := len(slice) - 1
	pivot := 0
	for ind1 <= ind2 {
		if slice[ind1] < slice[pivot] {
			swapElements(slice, pivot, ind1)
			ind1++
			pivot++
		}
		if slice[pivot] < slice[ind2] {
			ind2--
		}
		if slice[ind2] < slice[pivot] && slice[pivot] < slice[ind1] {
			swapElements(slice, ind1, ind2)
			swapElements(slice, pivot, ind1)
			ind1++
			ind2--
			pivot++
		}
	}

	// Recursion
	part1 := quickSortInPlace(slice[0 : pivot+1])
	part2 := quickSortInPlace(slice[pivot+1:])
	return slices.Concat(part1, part2)
}
