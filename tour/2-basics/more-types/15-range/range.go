package main

import "fmt"

var multiples = []int{0, 2, 4, 6, 8, 10}

func main() {
	for i, value := range multiples {
		fmt.Printf("2 x %d = %d\n", i, value)
	}

	exponents := make([]int, 10)

	for i := range exponents {
		exponents[i] = 1 << uint(i)
	}

	for _, value := range exponents {
		fmt.Printf("%d\n", value)
	}
}

/*

When iterating over a slice using range, each iteration returns an index and the element found at that index.

We can use _ to skip either the index or the value. The value can also be omitted altogether.

*/
