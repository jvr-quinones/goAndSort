package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Options struct {
	size   int
	sorter string
}

func parseArgs() Options {
	options := Options{}
	flag.IntVar(&options.size, "size", 1e3, "Sample size")
	flag.StringVar(&options.sorter, "sorter", "select", "Sorting algorithm")
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
	}

	if options.size <= 1 {
		fmt.Println("Size has to be greater than 1")
		os.Exit(2)
	}

	if !slices.Contains(algorithms, options.sorter) {
		fmt.Println("Unknown algortihm")
		os.Exit(2)
	}
}
