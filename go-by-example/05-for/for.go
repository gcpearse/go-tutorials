package main

import "fmt"

func main() {
	i := 0

	for i < 3 {
		fmt.Println("i =", i)
		i++
	}

	for j := 0; j < 3; j++ {
		fmt.Println("j =", j)
	}

	for k := range 3 {
		fmt.Println("k =", k)
	}

	for {
		fmt.Println("run")
		break
	}

	for n := range 6 {
		if n%2 != 0 {
			continue
		}

		fmt.Println("n =", n)
	}
}
