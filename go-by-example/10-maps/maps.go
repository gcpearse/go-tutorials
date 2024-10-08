package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 10
	m["k2"] = 20

	fmt.Println(m)

	v1 := m["k1"]

	fmt.Println(v1)

	v3 := m["k3"]

	fmt.Println(v3)

	fmt.Println(len(m))

	delete(m, "k2")

	fmt.Println(m)

	clear(m)

	fmt.Println(m)

	_, present := m["k2"]

	fmt.Println(present)

	n := map[string]int{"foo": 1, "bar": 2}

	fmt.Println(n)

	n2 := map[string]int{"foo": 1, "bar": 2}

	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}

/*

Use delete to remove specific key/value pairs and clear to remove all of them.

The optional second return value when retrieving a map value indicates whether the key is present in the map.

This can be useful for distinguishing between missing keys and keys with zero values.

*/
