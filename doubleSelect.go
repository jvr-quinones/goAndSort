package main

func doubleSelectSort(slice []int) (sorted []int) {
	sorted = make([]int, len(slice))
	copy(sorted, slice)
	length := len(sorted)

	for ind1 := 0; ind1 < length/2; ind1++ {
		smallest := ind1
		largest := ind1
		ind2 := ind1

		for ; ind2 <= length-ind1-1; ind2++ {
			if sorted[ind2] < sorted[smallest] {
				smallest = ind2
			}
			if sorted[ind2] > sorted[largest] {
				largest = ind2
			}
		}

		SwapElements(sorted, ind1, smallest)

		ind2-- // Added because the FOR loop does an extra "+1"
		if largest != ind2 && largest != ind1 {
			SwapElements(sorted, ind2, largest)

		} else if largest != ind2 && largest == ind1 {
			// Since the min swap will always occur first,
			// you have to follow the new position of the max value.
			// It only happens when the largest position is
			// in the same position as ind1
			SwapElements(sorted, ind2, smallest)

		}
	}

	return
}

func doubleSelectSortStable(slice []int) (sorted []int) {
	sorted = make([]int, len(slice))
	checked := make([]bool, len(slice))

	for ind1 := 0; ind1 < len(slice)/2; ind1++ {
		smallest := 0
		largest := 0
		readyToCheck := false

		for ind2 := 0; ind2 < len(slice); ind2++ {
			if !readyToCheck && !checked[ind2] {
				smallest = ind2
				largest = ind2
				readyToCheck = true
				continue
			}

			if readyToCheck && !checked[ind2] && slice[ind2] < slice[smallest] {
				smallest = ind2
			} else if readyToCheck && !checked[ind2] && slice[ind2] >= slice[largest] {
				largest = ind2
			}
		}

		sorted[ind1] = slice[smallest]
		sorted[len(slice)-ind1-1] = slice[largest]
		checked[smallest] = true
		checked[largest] = true
	}

	return sorted
}
