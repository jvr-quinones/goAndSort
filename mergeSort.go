package main

func mergeSortRecursive(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	arr1 := mergeSortRecursive(slice[0 : len(slice)/2])
	arr2 := mergeSortRecursive(slice[len(slice)/2:])
	return mergeArrays(arr1, arr2)
}

func mergeArrays(arr1 []int, arr2 []int) (merged []int) {
	merged = make([]int, len(arr1)+len(arr2))
	leftPtr := 0
	rightPtr := 0

	for ind1 := range merged {
		if leftPtr < len(arr1) && rightPtr < len(arr2) {
			if arr1[leftPtr] <= arr2[rightPtr] {
				merged[ind1] = arr1[leftPtr]
				leftPtr++
			} else {
				merged[ind1] = arr2[rightPtr]
				rightPtr++
			}

		} else if leftPtr >= len(arr1) {
			merged[ind1] = arr2[rightPtr]
			rightPtr++
		} else {
			merged[ind1] = arr1[leftPtr]
			leftPtr++
		}
	}

	return merged
}
