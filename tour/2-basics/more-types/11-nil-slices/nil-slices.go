package main

import "fmt"

func main() {
	var s []int

	fmt.Println(s, len(s), cap(s))
}

/*

The zero value of a slice is nil. A nil slice has a length and capacity of 0 and no underlying array.

*/
