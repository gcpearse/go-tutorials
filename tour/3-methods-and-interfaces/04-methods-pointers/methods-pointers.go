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

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)

	fmt.Println(v.Abs())
}

/*

Methods can be declared with pointer receivers for which the receiver type has the literal syntax *T for type T.

Here, Scale is defined on *Vertex and modifies the value to which the receiver points.

Pointer receivers are more common than value receivers, as a method often needs to modify its receiver.

With a value receiver, the Scale method operates on a copy of the original Vertex value.

The Scale method must have a pointer receiver to change the Vertex value declared in the main function.

*/
