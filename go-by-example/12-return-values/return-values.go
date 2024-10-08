package main

import "fmt"

func doubleTwoInts(x, y int) (int, int) {
	return x * 2, y * 2
}

func main() {
	a, b := doubleTwoInts(2, 5)

	fmt.Println(a)
	fmt.Println(b)

	_, c := doubleTwoInts(4, 10)

	fmt.Println(c)
}

/*

Go supports multiple return values.

We can use the blank identifier _ to return a subset of the return values.

*/
