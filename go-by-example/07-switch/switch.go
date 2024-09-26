package main

import (
	"fmt"
	"time"
)

func main() {
	isToggled := true

	switch isToggled {
	case true:
		fmt.Println("Toggle on")
	case false:
		fmt.Println("Toggle off")
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
