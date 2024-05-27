package main

import "fmt"

func PrintArray(slice []int) {
	for ind1 := range slice {
		fmt.Println(slice[ind1])
	}
}
