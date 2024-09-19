package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(str string, slice []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", str, len(slice), cap(slice), slice)
}

/*

The make function allocates a zeroed array and returns a slice that refers to that array.

A capacity can be specified by passing a third argument to make.

*/
