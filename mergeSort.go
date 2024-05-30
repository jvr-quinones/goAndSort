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

func mergeSortIterative(slice []int) (sorted []int) {
	sorted = make([]int, len(slice))
	intermediate := make([]int, len(slice))
	copy(sorted, slice)
	length := len(slice)

	for ind1 := 1; ind1 < len(slice); ind1 *= 2 {
		groupOffset := 0
		leftSize := ind1
		leftOffset := 0
		rightSize := 0
		rightOffset := 0

		if ind1 <= length/2 {
			rightSize = ind1
		} else {
			rightSize = length - leftSize
		}

		for ind2 := 0; ind2 < length; ind2++ {
			leftPtr := groupOffset*ind1 + leftOffset
			rightPtr := (groupOffset+1)*ind1 + rightOffset

			if leftOffset < leftSize && rightOffset < rightSize && leftPtr < length && rightPtr < length {
				if sorted[leftPtr] <= sorted[rightPtr] {
					intermediate[ind2] = sorted[leftPtr]
					leftOffset++
				} else {
					intermediate[ind2] = sorted[rightPtr]
					rightOffset++
				}

			} else if (rightOffset >= rightSize || rightPtr >= length) && leftPtr < length {
				intermediate[ind2] = sorted[leftPtr]
				leftOffset++

			} else if (leftOffset >= leftSize || leftPtr >= length) && rightPtr < length {
				intermediate[ind2] = sorted[rightPtr]
				rightOffset++

			} else {
				logger.Print("Something went wrong at:")
				logger.Printf("Group offset: %v", groupOffset)
				logger.Printf("Left pointer: %v", leftPtr)
				logger.Printf("Left offset: %v", leftOffset)
				logger.Printf("Right pointer: %v", rightPtr)
				logger.Printf("Right offset: %v", rightOffset)
				logger.Fatalln("Goodbye")
			}

			if leftOffset >= leftSize && rightOffset >= rightSize {
				leftOffset = 0
				rightOffset = 0
				groupOffset += 2
			}
		}

		copy(sorted, intermediate)
	}

	return sorted
}
