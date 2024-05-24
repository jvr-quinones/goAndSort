package main

import (
	"fmt"
	"strings"
)

var (
	array  []int
	sorted []int
)

func main() {
	fmt.Println("This program will eventually sort an array using different algorithms and sizes")
	options := parseArgs()

	array = make([]int, options.size)
	randomizeArray(array)
	fmt.Println("UNSORTED ARRAY")
	printArray(array)

	fmt.Println(strings.ToUpper(options.sorter), "SORT")

	switch options.sorter {
	case "select":
		sorted = selectSort(array)
	case "bubble":
		sorted = bubbleSort(array)
	}

	printArray(sorted)
}

func printArray(slice []int) {
	for _, val := range slice {
		fmt.Println(val)
	}
}
