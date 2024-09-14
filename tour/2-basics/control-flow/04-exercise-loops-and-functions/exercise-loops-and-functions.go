package main

import (
	"fmt"
)

// Given x, find the value of z for which z squared is closest to x.
func Sqrt(x float64) float64 {
	z := float64(1)

	for {
		if z == z-(z*z-x)/(2*z) {
			break
		} else {
			z -= (z*z - x) / (2 * z)
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

/*

An infinite loop can be created by omitting the loop condition.

*/
