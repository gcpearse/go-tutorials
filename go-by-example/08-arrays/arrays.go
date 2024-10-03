package main

import "fmt"

func main() {
	var a [5]int

	fmt.Println(a)

	a[4] = 100

	fmt.Println(a)
	fmt.Println(a[4])
	fmt.Println(len(a))

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println(b)

	c := [...]int{1, 2, 3, 4, 5}

	fmt.Println(c)

	d := [...]int{10, 3: 40, 50}

	fmt.Println(d)

	var nested [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			nested[i][j] = i + j
		}
	}

	fmt.Println(nested)

	nested = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println(nested)
}

/*

We can use ... syntax to have the compiler count the number of elements for us.

When an index is specified with :, the elements in between will be zeroed.

*/
