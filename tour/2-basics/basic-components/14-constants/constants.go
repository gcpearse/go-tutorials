package main

import "fmt"

const Pi = 3.14

func main() {
	fmt.Println("Happy", Pi, "Day!")

	const World = "Earth"
	fmt.Println("Hello,", World)

	const IsTrue = true
	fmt.Println("True or false? Go is awesome:", IsTrue)
}

/*

Constants cannot be declared using the := syntax.

*/
