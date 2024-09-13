package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}

/*

With this code we have:

- declared a main package
- imported the fmt and rsc.io/quote packages
- written a main function to print messages to the console

*/
