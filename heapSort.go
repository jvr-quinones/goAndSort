package main

func heapSort(slice []int) (sorted []int) {
	sorted = make([]int, len(slice))
	copy(sorted, slice)
	ind1 := 0

	for ind1 = 1; ind1 < len(sorted); ind1++ {
		for ind2 := ind1; ind2 > 0; {
			parent := (ind2 - 1) / 2
			if sorted[ind2] > sorted[parent] {
				swapElements(sorted, ind2, parent)
				ind2 = parent
			} else {
				break
			}
		}
	}

	for ind1 = len(sorted) - 1; ind1 > 0; ind1-- {
		swapElements(sorted, 0, ind1)
		pushValueDown(sorted, 0, ind1)
	}

	if sorted[0] > sorted[1] {
		swapElements(sorted, 0, 1)
	}
	return sorted
}

func pushValueDown(slice []int, current int, hiLim int) {
	for ind1 := current; ind1 < hiLim/2; {
		leftChild := 2*ind1 + 1
		rightChild := leftChild + 1

		if slice[ind1] < slice[leftChild] && slice[leftChild] > slice[rightChild] {
			swapElements(slice, ind1, leftChild)
			ind1 = leftChild

		} else if slice[ind1] < slice[rightChild] && rightChild < hiLim {
			swapElements(slice, ind1, rightChild)
			ind1 = rightChild

		} else {
			break
		}
	}
}
