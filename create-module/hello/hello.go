package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"James", "Danny", "Rob"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}

/*

The log package is configured to print the command name ("greeting: ") at the beginning of log messages.

If there is an error, the log package's Fatal function prints the error and stops the program.

*/
