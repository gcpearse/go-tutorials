package main

import "fmt"

func main() {
	primes := [5]int{2, 3, 5, 7, 11}

	slice := primes[1:4]

	fmt.Println(slice)
}

/*

A slice is formed by specifying a lower bound and a higher bound separated by a colon.

*/
