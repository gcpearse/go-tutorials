package main

import "fmt"

type Vertex struct {
	Latitude, Longitude float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}

/*

A map literal is similar to a struct literal but the keys are required.

*/
