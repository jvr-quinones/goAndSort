package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

type Options struct {
	size   int
	sorter string
}

var (
	array  []int
	sorted []int
)

func main() {
	options := parseArgs()
	fmt.Println("This program will eventually sort an array using different algorithms and sizes")

	array = make([]int, options.size)
	RandomizeArray(array)
	fmt.Println("UNSORTED ARRAY")
	PrintArray(array)
	fmt.Println(strings.ToUpper(options.sorter), "SORT")

	timeInit := time.Now()
	switch options.sorter {
	case "binary-insert":
		sorted = binaryInsertSort(array)
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
	case "shaker":
		sorted = shakerSort(array)
	}
	duration := time.Since(timeInit)

	PrintArray(sorted)
	fmt.Println("Duration:", duration.Microseconds(), "us")

	isSorted := "no"
	if slices.IsSorted(sorted) {
		isSorted = "yes"
	}
	fmt.Println("Is it sorted??", isSorted)
}

func parseArgs() Options {
	options := Options{}
	flag.IntVar(&options.size, "size", 1e3, "Sample size")
	flag.StringVar(&options.sorter, "sorter", "binary-insert", "Sorting algorithm")
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
		"exchange",
		"insert",
		"select",
		"shaker",
	}

	logger := log.New(os.Stderr, "goAndSortError: ", 0)
	if options.size <= 1 {
		logger.Fatal("Size has to be greater than 1")
	}

	if !slices.Contains(algorithms, options.sorter) {
		logger.Fatalf("Unknown algorithm %q", options.sorter)
	}
}
