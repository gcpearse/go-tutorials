package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	sum = 1

	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

/*

A basic for loop is made up of an init statement, condition expression, and post statement.

The init and post statements are optional. A for loop without an init statement and post statement works like a while loop.

*/
