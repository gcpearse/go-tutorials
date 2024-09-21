package main

import "fmt"

type Vertex struct {
	Latitude, Longitude float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])
}

/*

A map contains key-value pairs.

The make function initialises a map of the given type.

The zero value of a map is nil.

*/
