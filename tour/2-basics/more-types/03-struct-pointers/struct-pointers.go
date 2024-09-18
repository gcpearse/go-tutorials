package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	p := &v
	p.X = 3

	fmt.Println(v)
}

/*

Go allows us to write p.X rather than (*p).X, avoiding the explicit dereference.

*/
