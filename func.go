package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	name := "max"
	age := 101
	var message string

	info, access, err := enterTheClub(age)
	if err != nil {
		// enter error and quite code
		log.Fatal(err)
	}

	if access {
		message = sayHello(name, age) + info
	} else {
		message = info
	}

	printMessage(message)
}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Hello,  %s! You %d years old, ", name, age)
}

// множественные возращаемые значения
func enterTheClub(age int) (string, bool, error) {
	if age >= 18 && age < 100 {
		return "Please enter", true, nil
	}

	if age >= 100 {
		return "", false, errors.New("you are dead man")
	}

	return "Access denied", false, nil
}
