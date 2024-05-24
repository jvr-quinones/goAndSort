package main

func selectAndSort(slice []int) (sorted []int) {
	sorted = make([]int, len(slice))
	copy(sorted, slice)

	for ind1 := 0; ind1 < (len(sorted) - 1); ind1++ {
		smallest := ind1
		for ind2 := ind1 + 1; ind2 < (len(sorted)); ind2++ {
			if sorted[ind2] < sorted[smallest] {
				smallest = ind2
			}
		}
		if smallest != ind1 {
			tempVal := sorted[ind1]
			sorted[ind1] = sorted[smallest]
			sorted[smallest] = tempVal
		}
	}

	return sorted
}
