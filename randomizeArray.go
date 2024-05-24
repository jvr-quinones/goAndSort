package main

import (
	"math/rand"
)

func randomizeArray(slice []int) {
	for ind := range slice {
		slice[ind] = rand.Intn(len(slice) * 1e2)
	}
}
