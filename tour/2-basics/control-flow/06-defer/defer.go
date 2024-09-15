package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("Hello")

	fmt.Println("Counting...")

	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}

/*

Deferred function calls are pushed onto a stack and executed following the last-in-first-out principle.

*/
