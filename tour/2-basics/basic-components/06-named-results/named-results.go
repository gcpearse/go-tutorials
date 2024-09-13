package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 1 / 2
	y = sum - x
	return
}

func main() {
	fmt.Println(split(10))
}

/*

A return statement without arguments returns the named return values. This is called a naked return.

*/
