package main

import "fmt"

func addTwoInts(a int, b int) int {
	return a + b
}

func addThreeInts(a, b, c int) int {
	return a + b + c
}

func main() {
	res := addTwoInts(1, 2)
	fmt.Println(res)

	res = addThreeInts(1, 2, 3)
	fmt.Println(res)
}
