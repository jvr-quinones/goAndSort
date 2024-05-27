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
