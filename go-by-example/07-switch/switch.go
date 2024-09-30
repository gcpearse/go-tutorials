package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	getType := func(i interface{}) {
		switch t := i.(type) {
		case string:
			fmt.Println("This is a string")
		case int:
			fmt.Println("This is an integer")
		case bool:
			fmt.Println("This is a Boolean value")
		default:
			fmt.Printf("Type: %T\n", t)
		}
	}

	getType("hello")
	getType(1)
	getType(true)
	getType(1.1)
}

/*

A switch statement without an expression is an alternative way of expressing if/else logic.

A type switch can be used to identify the type of an interface value.

*/
