package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Options struct {
	size   int
	sorter string
}

var (
	array  []int
	logger *log.Logger
	sorted []int
)

func main() {
	logger = log.New(os.Stderr, "goAndSort: ", 0)
	options := parseArgs()
	fmt.Println("This program will eventually sort an array using different algorithms and sizes")

	array = make([]int, options.size)
	randomizeArray(array)
	fmt.Println("UNSORTED ARRAY")
	printArray(array)
	fmt.Println(strings.ToUpper(strings.ReplaceAll(options.sorter, "-", " ")), "SORT")

	switch options.sorter {
	case "binary-insert":
		sorted = binaryInsertSort(array)
	case "bubble":
		sorted = bubbleSort(array)
	case "double-select":
		sorted = doubleSelectSort(array)
	case "double-select-stable":
		sorted = doubleSelectSortStable(array)
	case "exchange":
		sorted = exchangeSort(array)
	case "heap":
		sorted = heapSort(array)
	case "insert":
		sorted = insertSort(array)
	case "merge-recursive":
		sorted = mergeSortRecursive(array)
	case "merge-iterative":
		sorted = mergeSortIterative(array)
	case "select":
		sorted = selectSort(array)
	case "select-stable":
		sorted = selectSortStable(array)
	case "shaker":
		sorted = shakerSort(array)
	default:
		logger.Fatalf("No handler for sorter %q", options.sorter)
	}
	printArray(sorted)

	isSorted := "no"
	if slices.IsSorted(sorted) {
		isSorted = "yes"
	}
	fmt.Println("Is it sorted??", isSorted)
}

func parseArgs() Options {
	options := Options{}
	flag.IntVar(&options.size, "size", 1e3, "Sample size")
	flag.StringVar(&options.sorter, "sorter", "merge-iterative", "Sorting algorithm")
	flag.Parse()
	options.sorter = strings.ToLower(options.sorter)

	checkArgs(&options)
	return options
}

func checkArgs(options *Options) {
	algorithms := []string{
		"binary-insert",
		"bubble",
		"double-select",
		"double-select-stable",
		"exchange",
		"heap",
		"insert",
		"merge-iterative",
		"merge-recursive",
		"select",
		"select-stable",
		"shaker",
	}

	if options.size <= 1 {
		logger.Fatal("Size has to be greater than 1")
	}

	if !slices.Contains(algorithms, options.sorter) {
		logger.Fatalf("Unknown algorithm %q", options.sorter)
	}
}
