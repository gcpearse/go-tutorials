package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 10
	fmt.Println("The answer is:", m["Answer"])

	m["Answer"] = 15
	fmt.Println("The answer is:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The answer is:", m["Answer"])

	value, isPresent := m["Answer"]
	fmt.Println("The answer is:", value, "Is the key present?", isPresent)
}

/*

We can test whether or not a key is present with a two-value assignment.

*/
