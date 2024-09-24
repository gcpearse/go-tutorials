package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

/*

Constant expressions perform arithmetic with arbitrary precision.

A numeric constant has no type until given one. For example, we can use an explicit conversion.

A numeric constant can also be given a type by using it in a context that requires a specific type. For example, math.Sin expects a float64.

*/
