package main

func countingSort(slice []int) (sorted []int) {
	max := 0
	ind1 := 0

	for ind1 = range slice {
		if slice[ind1] > max {
			max = slice[ind1]
		}
	}

	counter := make([]int, max+1)
	sorted = make([]int, len(slice))

	for _, val := range slice {
		counter[val] += 1
	}

	for ind1 = 0; ind1 < len(counter)-1; ind1++ {
		counter[ind1+1] += counter[ind1]
	}

	for ind1 = len(slice) - 1; ind1 >= 0; ind1-- {
		val := slice[ind1]
		counter[val]--
		pos := counter[val]
		sorted[pos] = val
	}

	return sorted
}
