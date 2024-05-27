package main

import "fmt"

func shakerSort(slice []int) (sorted []int) {
	sorted = make([]int, len(slice))
	copy(sorted, slice)

	for ind1 := 0; ind1 < len(sorted)/2; ind1++ {
		dirRev := false
		limit := len(slice) - ind1 - 1

		for ind2 := ind1; ind2 > ind1 || !dirRev; {
			if !dirRev && sorted[ind2] > sorted[ind2+1] {
				err := SwapElements(sorted, ind2, ind2+1)
				if err != nil {
					fmt.Println(err.Error())
				}
			} else if dirRev && sorted[ind2] < sorted[ind2-1] {
				err := SwapElements(sorted, ind2-1, ind2)
				if err != nil {
					fmt.Println(err.Error())
				}
			}

			if dirRev {
				ind2--
			} else {
				ind2++
			}

			if ind2 == limit {
				dirRev = true
			}
		}
	}

	return sorted
}
