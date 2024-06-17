package main

import "fmt"

func rawPrintArray(slice []int) {
	for ind1 := range slice {
		fmt.Println(slice[ind1])
	}
}

func csvPrintArray(slices ...[]int) {
	arrayLengths := make([]int, len(slices))
	rows := 0
	maxLength := 0

	for rows = range slices {
		arrayLengths[rows] = len(slices[rows])
		if arrayLengths[rows] > maxLength {
			maxLength = arrayLengths[rows]
		}
	}

	fmt.Println("UNSORTED,SORTED")
	for rows = 0; rows < maxLength; rows++ {
		var csvLine string

		for cols := range slices {
			if cols > 0 {
				csvLine += ","
			}

			if cols < len(slices[cols]) {
				csvLine += fmt.Sprintf("%d", slices[cols][rows])
			}
		}

		fmt.Println(csvLine)
	}
}
