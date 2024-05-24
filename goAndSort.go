package main

import (
	"fmt"
)

var (
	array [1e3]int
)

func main() {
	fmt.Println("This program will eventually sort an array")
	randomizeArray(array[:])
	fmt.Println("UNSORTED ARRAY")
	printArray(array[:])
}

func printArray(slice []int) {
	for _, val := range slice {
		fmt.Println(val)
	}
}
