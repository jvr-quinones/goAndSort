package main

import "fmt"

func printArray(slice []int) {
	for ind1 := range slice {
		fmt.Println(slice[ind1])
	}
}
