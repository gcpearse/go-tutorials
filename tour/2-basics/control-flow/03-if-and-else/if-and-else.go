package main

import (
	"fmt"
	"math"
)

func pow(x, exponent, limit float64) float64 {
	if v := math.Pow(x, exponent); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
	}

	return limit
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

/*

An if statement can begin with a short statement to be executed before the condition.

*/
