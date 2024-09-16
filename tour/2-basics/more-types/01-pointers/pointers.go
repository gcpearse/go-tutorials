package main

import "fmt"

func main() {
	i, j := 10, 20

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 15         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 4    // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

/*

A pointer stores the memory address of a value.

The & operator generates a pointer to its operand.

The * operator denotes the pointer's underlying value.

*/
