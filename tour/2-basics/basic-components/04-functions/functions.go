package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(10, 20))
}

/*

When two or more consecutive function parameters share a type, only the last parameter needs to be annotated.

*/
