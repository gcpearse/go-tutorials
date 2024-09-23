package main

import "fmt"

var i, j int = 1, 2

func main() {
	var a, b, c = true, false, "no!"

	fmt.Println(i, j, a, b, c)
}

/*

When an initialiser is present, the variable takes the type of the initialiser.

*/
