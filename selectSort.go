package main

func selectSort(slice []int) (sorted []int) {
	sorted = make([]int, len(slice))
	copy(sorted, slice)

	for ind1 := 0; ind1 < (len(sorted) - 1); ind1++ {
		smallest := ind1
		for ind2 := ind1 + 1; ind2 < (len(sorted)); ind2++ {
			if sorted[ind2] < sorted[smallest] {
				smallest = ind2
			}
		}

		swapElements(sorted, ind1, smallest)
	}

	return sorted
}

func selectSortStable(slice []int) (sorted []int) {
	checked := make([]bool, len(slice))
	sorted = make([]int, len(slice))

	for ind1 := 0; ind1 < len(slice); ind1++ {
		smallest := 0
		readyToCheck := false

		for ind2 := 0; ind2 < len(slice); ind2++ {
			if !readyToCheck && !checked[ind2] {
				smallest = ind2
				readyToCheck = true
				continue
			}

			if readyToCheck && !checked[ind2] && slice[ind2] < slice[smallest] {
				smallest = ind2
			}
		}

		sorted[ind1] = slice[smallest]
		checked[smallest] = true
	}

	return sorted
}
