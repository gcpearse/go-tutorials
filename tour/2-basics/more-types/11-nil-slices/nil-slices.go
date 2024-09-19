package main

import "fmt"

func main() {
	var slice []int

	fmt.Println(slice, len(slice), cap(slice))
}

/*

The zero value of a slice is nil. A nil slice has a length and capacity of 0 and no underlying array.

*/
