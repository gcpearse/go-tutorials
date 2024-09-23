package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypotenuse := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypotenuse(5, 12))   // Find the length of the hypotenuse when a = 5 and b = 12.
	fmt.Println(compute(hypotenuse)) // Pass hypotenuse to compute to find the length of the hypotenuse when a = 3 and b = 4.
	fmt.Println(compute(math.Pow))   // Pass math.Pow to compute to find the value of 3**4.
}
