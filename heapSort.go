package main

func heapSort(slice []int) (sorted []int) {
	sorted = make([]int, len(slice))
	copy(sorted, slice)
	ind1 := 0

	for ind1 = 1; ind1 < len(sorted); ind1++ {
		pushValueUp(sorted, ind1)
		}

	for ind1 = len(sorted) - 1; ind1 > 0; ind1-- {
		SwapElements(sorted, 0, ind1)
		pushValueDown(sorted, 0, ind1)
	}

	if sorted[0] > sorted[1] {
		SwapElements(sorted, 0, 1)
	}
	return sorted
}

func pushValueUp(slice []int, current int) {
	for ind1 := current; ind1 > 0; {
		parent := (ind1 - 1) / 2

		if slice[ind1] > slice[parent] {
			SwapElements(slice, ind1, parent)
			ind1 = parent

		} else {
			break
		}
	}
}

func pushValueDown(slice []int, current int, hiLim int) {
	for ind1 := current; ind1 < hiLim/2; {
		leftChild := 2*ind1 + 1
		rightChild := 2*ind1 + 2

		if slice[ind1] < slice[leftChild] && slice[leftChild] > slice[rightChild] {
			SwapElements(slice, ind1, leftChild)
			ind1 = leftChild

		} else if slice[ind1] < slice[rightChild] && rightChild < hiLim {
			SwapElements(slice, ind1, rightChild)
			ind1 = rightChild

		} else {
			break
		}
	}
}
