package main

import "fmt"

func main() {
	n := 12

	if n%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 0; num < 0 {
		fmt.Println(num, "is smaller than zero")
	} else if num == 0 {
		fmt.Println(num, "is zero")
	} else {
		fmt.Println(num, "is greater than zero")
	}
}
