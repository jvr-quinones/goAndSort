package main

func gnomeSort(slice []int) (sorted []int) {
	sorted = make([]int, len(slice))
	copy(sorted, slice)

	pos := 0
	moveUp := true
	for pos < len(sorted)-1 {
		if sorted[pos] > sorted[pos+1] {
			swapElements(sorted, pos, pos+1)
			moveUp = false
		} else if sorted[pos] < sorted[pos+1] {
			moveUp = true
		}

		if moveUp {
			pos++
		} else if pos > 0 {
			pos--
		} else {
			moveUp = true
			pos++
		}

	}

	return sorted
}
