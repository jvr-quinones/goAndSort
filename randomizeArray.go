package main

import (
	"math/rand"
)

func RandomizeArray(slice []int) {
	for ind := range slice {
		slice[ind] = rand.Intn(len(slice) * 1e2)
	}
}
