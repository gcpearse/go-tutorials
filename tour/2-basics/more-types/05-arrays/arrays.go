package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes)
}
