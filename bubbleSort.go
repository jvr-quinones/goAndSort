package main

func bubbleSort(slice []int) (sorted []int) {
	sorted = make([]int, len(slice))
	copy(sorted, slice)

	for ind1 := 0; ind1 < len(sorted)-1; ind1++ {
		for ind2 := 0; ind2 < len(sorted)-ind1-1; ind2++ {
			if sorted[ind2] > sorted[ind2+1] {
				swapElements(sorted, ind2, ind2+1)
			}
		}
	}

	return sorted
}
