package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New(("empty name"))
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hello, %v!",
		"Great to see you, %v!",
		"%v! Hail and well met!",
	}

	return formats[rand.Intn(len(formats))]
}

/*

A Go function can return multiple values. The nil value returned by Hello means there was no error.

The for loop in Hellos uses a blank identifier since the index is not needed.

In Go, a map (set of key value pairs) is initialised with the make function.

The Hello and Hellos functions are exported. Since the randomFormat function starts with a lowercase letter, it can only be accessed by the code in this package.

The formats variable is a slice: an array with a size that changes dynamically.

*/
