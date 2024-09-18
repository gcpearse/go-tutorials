package main

import "fmt"

func main() {
	var array [2]string
	array[0] = "Hello"
	array[1] = "world"

	fmt.Println(array[0], array[1])
	fmt.Println(array)

	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes)
}
