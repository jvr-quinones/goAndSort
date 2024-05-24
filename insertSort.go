package main

func insertSort(slice []int) (sorted []int) {
	sorted = make([]int, len(slice))
	copy(sorted, slice)

	for ind1 := 1; ind1 < len(slice); ind1++ {
		for ind2 := ind1; ind2 > 0; ind2-- {
			if sorted[ind2-1] <= sorted[ind2] {
				break
			}
			tempVal := sorted[ind2-1]
			sorted[ind2-1] = sorted[ind2]
			sorted[ind2] = tempVal
		}
	}

	return sorted
}
