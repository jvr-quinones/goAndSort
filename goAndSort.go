package main

import (
	"fmt"
	"slices"
	"strings"
	"time"
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

	timeInit := time.Now()
	switch options.sorter {
	case "bubble":
		sorted = bubbleSort(array)
	case "double-select":
		sorted = doubleSelectSort(array)
	case "exchange":
		sorted = exchangeSort(array)
	case "insert":
		sorted = insertSort(array)
	case "select":
		sorted = selectSort(array)
	}

	duration := time.Since(timeInit)
	printArray(sorted)
	fmt.Println("Duration:", duration.Microseconds(), "us")

	isSorted := "no"
	if slices.IsSorted(sorted) {
		isSorted = "yes"
	}
	fmt.Println("Is it sorted??", isSorted)
}

func printArray(slice []int) {
	for _, val := range slice {
		fmt.Println(val)
	}
}
