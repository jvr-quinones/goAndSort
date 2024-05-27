package main

import "errors"

func SwapElements(slice []int, a int, b int) error {
	length := len(slice)

	if a >= length {
		return errors.New("parameter a is greater than length of array")
	} else if b >= length {
		return errors.New("parameter a is greater than length of array")
	} else if a != b {
		tempVal := slice[a]
		slice[a] = slice[b]
		slice[b] = tempVal
	}
	return nil
}
