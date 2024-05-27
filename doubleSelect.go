package main

func doubleSelectSort(slice []int) (sorted []int) {
	sorted = make([]int, len(slice))
	copy(sorted, slice)
	length := len(sorted)

	for ind1 := 0; ind1 < length/2; ind1++ {
		smallest := ind1
		largest := ind1
		tempVal := 0

		ind2 := ind1
		for ; ind2 <= length-ind1-1; ind2++ {
			if sorted[ind2] < sorted[smallest] {
				smallest = ind2
			}
			if sorted[ind2] > sorted[largest] {
				largest = ind2
			}
		}

		if smallest != ind1 {
			tempVal = sorted[ind1]
			sorted[ind1] = sorted[smallest]
			sorted[smallest] = tempVal
		}

		ind2-- // Added because the FOR loop does an extra "+1"
		if largest != ind2 && largest != ind1 {
			tempVal = sorted[ind2]
			sorted[ind2] = sorted[largest]
			sorted[largest] = tempVal

		} else if largest != ind2 && largest == ind1 {
			// Since the min swap will always occur first,
			// you have to follow the new position of the max value.
			// It only happens when the largest position is
			// in the same position as ind1

			tempVal = sorted[ind2]
			sorted[ind2] = sorted[smallest]
			sorted[smallest] = tempVal
		}
	}

	return
}
