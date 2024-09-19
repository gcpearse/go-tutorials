package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11}
	printSlice(primes)

	primes = primes[:0]
	printSlice(primes)

	primes = primes[:4]
	printSlice(primes)

	primes = primes[2:]
	printSlice(primes)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
