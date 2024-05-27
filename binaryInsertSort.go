package main

func binaryInsertSort(slice []int) (sorted []int) {
	sorted = make([]int, len(slice))
	copy(sorted, slice)

	for ind1 := 1; ind1 < len(slice); ind1++ {
		// Insert binary search here
		posLeft := 0
		posRight := ind1

		for posLeft != posRight {
			posMid := (posLeft + posRight) / 2
			if sorted[posMid] > sorted[ind1] {
				posRight = posMid
			} else {
				posLeft = posMid + 1
			}
		}

		for ind2 := ind1; ind2 > posRight; ind2-- {
			tempVal := sorted[ind2-1]
			sorted[ind2-1] = sorted[ind2]
			sorted[ind2] = tempVal
		}
	}

	return sorted
}

// Returns the number of elements that are greater or equal to val
func BinarySeachRightmost(slice []int, val int) int {
	posLeft := 0
	posRight := len(slice) - 1

	for posLeft != posRight {
		posMid := (posLeft + posRight) / 2
		if sorted[posMid] > val {
			posRight = posMid
		} else {
			posLeft = posMid + 1
		}
	}

	return posRight
}

// Returns the number of elements that are less than val
func BinarySeachLeftmost(slice []int, val int) int {
	posLeft := 0
	posRight := len(slice) - 1

	for posLeft != posRight {
		posMid := (posLeft + posRight) / 2
		if sorted[posMid] > val {
			posRight = posMid + 1
		} else {
			posLeft = posMid
		}
	}

	return posLeft
}
