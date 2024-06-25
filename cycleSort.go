package main

func cycleSort(slice []int) (sorted []int) {
	sorted = make([]int, len(slice))
	copy(sorted, slice)

	for ind1 := 0; ind1 < len(sorted)-1; ind1++ {
		counter := 0
		valCurrent := sorted[ind1]
		firstScan := true
		isSorted := true

		for counter > 0 || firstScan {
			counter = 0

			for ind2 := ind1 + 1; ind2 < len(sorted); ind2++ {
				if valCurrent > sorted[ind2] {
					counter++
				}
				isSorted = isSorted && sorted[ind2-1] < sorted[ind2]
			}

			for counter > 0 && sorted[counter+ind1] == valCurrent {
				counter++
			}

			if counter >= len(sorted) {
				logger.Fatalf("position out of bounds with ind1=%d and counter=%d", ind1, counter)
			} else if counter > 0 || !firstScan {
				pos := counter + ind1
				valNew := sorted[pos]
				sorted[pos] = valCurrent
				valCurrent = valNew
			}

			firstScan = false
		}

		if isSorted {
			break
		}
	}

	return sorted
}
