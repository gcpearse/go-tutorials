package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	bools := []bool{false, true, true, false, true}
	fmt.Println(bools)

	structs := []struct {
		i int
		b bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
	}

	fmt.Println(structs)
}

/*

A slice literal is like an array literal without a specified length.

[]int{1, 2, 3, 4, 5} creates an array of length 5 and builds a slice that references it.

*/
