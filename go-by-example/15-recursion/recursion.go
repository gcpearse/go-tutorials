package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(5))

	var fibonacci func(n int) int

	fibonacci = func(n int) int {
		if n < 2 {
			return n
		}

		return fibonacci(n-1) + fibonacci(n-2)
	}

	fmt.Println(fibonacci(10))
}

/*

An anonymous function can be recursive if it is explicitly declared with var before being defined.

*/
