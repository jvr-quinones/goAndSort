package main

func doubleSelectSort(slice []int) (sorted []int) {
	sorted = make([]int, len(slice))
	copy(sorted, slice)

	length := len(sorted)
	tempVal := 0
	ind1 := 0
	ind2 := length - 1
	for ind1 < ind2 {
		smallest := ind1
		largest := ind1

		for ind3 := ind1 + 1; ind3 <= ind2; ind3++ {
			if sorted[ind3] < sorted[smallest] {
				smallest = ind3
			}
			if sorted[ind3] > sorted[largest] {
				largest = ind3
			}
		}

		if smallest != ind1 {
			tempVal = sorted[ind1]
			sorted[ind1] = sorted[smallest]
			sorted[smallest] = tempVal
		}

		if largest != ind2 && largest != ind1 {
			tempVal = sorted[ind2]
			sorted[ind2] = sorted[largest]
			sorted[largest] = tempVal

			// This case occurs because the max value used to be in ind1
			// but it was moved due to the swap with the min
		} else if largest != ind2 && largest == ind1 {
			tempVal = sorted[ind2]
			sorted[ind2] = sorted[smallest]
			sorted[smallest] = tempVal
		}

		ind1++
		ind2--
	}

	return
}
