package main

import "fmt"

func sum(nums ...int) int {
	fmt.Print(nums, " ")

	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))

	nums := []int{1, 2, 3, 4}

	fmt.Println(sum(nums...))

	fmt.Println(sum(1, 2, 3, 4, 5))
}
