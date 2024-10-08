package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string

	fmt.Println(s, s == nil, len(s) == 0)

	s = make([]string, 3)

	fmt.Println(s, len(s), cap(s))

	s[0], s[1], s[2] = "a", "b", "c"

	fmt.Println(s)
	fmt.Println(s[2])
	fmt.Println(len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println(c)

	l := s[2:5]

	fmt.Println(l)

	l = s[:5]

	fmt.Println(l)

	l = s[2:]

	fmt.Println(l)

	t := []string{"g", "h", "i"}
	fmt.Println(t)

	t2 := []string{"g", "h", "i"}

	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	nested := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLength := i + 1
		nested[i] = make([]int, innerLength)

		for j := 0; j < innerLength; j++ {
			nested[i][j] = i + j
		}
	}

	fmt.Println(nested)
}

/*

An uninitialised slice equals nil and has a length of 0.

Use make to create an empty slice with a set length.

By default, a new slice's capacity is equal to its length, but we can also explicitly pass capacity as an additional parameter to make.

Unlike multidimensional arrays, multidimensional slices can contain inner slices of varying lengths.

*/
