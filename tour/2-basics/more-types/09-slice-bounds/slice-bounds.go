package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11}

	slice1 := primes[1:4]
	fmt.Println(slice1)

	slice2 := primes[:2]
	fmt.Println(slice2)

	slice3 := primes[1:]
	fmt.Println(slice3)

	slice4 := primes[:]
	fmt.Println(slice4)
}
