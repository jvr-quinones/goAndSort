package main

// Returns the number of elements that are greater or equal to val
func BinarySeachRightmost(slice []int, val int) int {
	posLeft := 0
	posRight := len(slice) - 1

	for posLeft != posRight {
		posMid := (posLeft + posRight) / 2
		if slice[posMid] > val {
			posRight = posMid
		} else {
			posLeft = posMid + 1
		}
	}

	return posRight
}

// Returns the number of elements that are less than val
func BinarySeachLeftmost(slice []int, val int) int {
	posLeft := 0
	posRight := len(slice) - 1

	for posLeft != posRight {
		posMid := (posLeft + posRight) / 2
		if slice[posMid] > val {
			posRight = posMid + 1
		} else {
			posLeft = posMid
		}
	}

	return posLeft
}
