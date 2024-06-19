package main

import (
	"flag"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

type Options struct {
	// Array size
	size int

	// Algorithm used
	sorter string

	// If true, print out in json format, print in csv format otherwise
	// json bool

	// Print begin and end timestamps (in microseconds maybe)
	timestamps bool
}

var (
	logger *log.Logger
)

func main() {
	var (
		timeStart int64
		timeEnd   int64
		sorted    []int
	)

	logger = log.New(os.Stderr, "goAndSort: ", 0)
	options := parseArgs()

	timeStart = time.Now().UnixMicro()
	unsorted := make([]int, options.size)
	randomizeArray(unsorted)

	switch options.sorter {
	case "binary-insert":
		sorted = binaryInsertSort(unsorted)
	case "bubble":
		sorted = bubbleSort(unsorted)
	case "counting":
		sorted = countingSort(unsorted)
	case "double-select":
		sorted = doubleSelectSort(unsorted)
	case "double-select-stable":
		sorted = doubleSelectSortStable(unsorted)
	case "exchange":
		sorted = exchangeSort(unsorted)
	case "gnome":
		sorted = gnomeSort(unsorted)
	case "heap":
		sorted = heapSort(unsorted)
	case "insert":
		sorted = insertSort(unsorted)
	case "merge-recursive":
		sorted = mergeSortRecursive(unsorted)
	case "merge-iterative":
		sorted = mergeSortIterative(unsorted)
	case "quick":
		sorted = quickSort(unsorted)
	case "radix-sort-base2-lsd":
		sorted = radixSortLSD(unsorted, 2)
	case "radix-sort-base10-lsd":
		sorted = radixSortLSD(unsorted, 10)
	case "radix-sort-base10-msd":
		sorted = radixSortMSD(unsorted, 10)
	case "select":
		sorted = selectSort(unsorted)
	case "select-stable":
		sorted = selectSortStable(unsorted)
	case "shaker":
		sorted = shakerSort(unsorted)
	default:
		logger.Fatalf("No handler for sorter %q", options.sorter)
	}

	csvPrintArray(unsorted, sorted)
	isSorted := slices.IsSorted(sorted)
	timeEnd = time.Now().UnixMicro()

	if options.timestamps {
		logger.Printf("Start time: %v", timeStart)
		logger.Printf("End time: %v", timeEnd)
	}

	if isSorted && len(sorted) > 0 {
		logger.Println("Array was sorted")
	} else {
		logger.Fatal("Array was not sorted")
	}
}

func parseArgs() Options {
	options := Options{}
	flag.IntVar(&options.size, "size", 1e3, "Sample size")
	flag.StringVar(&options.sorter, "sorter", "merge-iterative", "Sorting algorithm")
	// flag.BoolVar(&options.json, "json", false, "If true, output is in json format, output is csv format otherwise")
	flag.BoolVar(&options.timestamps, "timestamps", false, "Print the start and end timestamp in unix time")
	flag.Parse()
	options.sorter = strings.ToLower(options.sorter)

	checkArgs(&options)
	return options
}

func checkArgs(options *Options) {
	algorithms := []string{
		"binary-insert",
		"bubble",
		"counting",
		"double-select",
		"double-select-stable",
		"exchange",
		"gnome",
		"heap",
		"insert",
		"merge-iterative",
		"merge-recursive",
		"quick",
		"radix-sort-base2-lsd",
		"radix-sort-base10-lsd",
		"radix-sort-base10-msd",
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
