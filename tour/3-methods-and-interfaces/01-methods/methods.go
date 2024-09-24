package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}

	fmt.Println(v.Abs())
}

/*

Go does not have classes but we can define methods on types.

A method takes a receiver argument between the func keyword and the method name.

In this example, the Abs method has a receiver - v - of type Vertex.

*/
