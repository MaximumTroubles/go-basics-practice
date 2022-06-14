package main

import (
	"errors"
	"fmt"
)

func main() {
	switchCase("Helo")
}

func switchCase(message string) {
	switch message {
	case "Hello":
		fmt.Printf("Wrong")
	case "HeLLO":
		fmt.Printf("two big LL")
	case "hello":
		fmt.Printf("All good")
	default:
		fmt.Println(errors.New("made mistake in word => " + message))
	}
}
