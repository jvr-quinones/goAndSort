package main

import (
	"fmt"
)

var (
	array  [1e1]int
	sorted []int
)

func main() {
	fmt.Println("This program will eventually sort an array")
	randomizeArray(array[:])
	fmt.Println("UNSORTED ARRAY")
	printArray(array[:])

	sorted = selectAndSort(array[:])
	fmt.Println("SELECT SORT")
	printArray(sorted)
}

func printArray(slice []int) {
	for _, val := range slice {
		fmt.Println(val)
	}
}
