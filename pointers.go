package main

import "fmt"

func main() {
	// cell im memory 0xc000010250
	message := "Golang "
	// operation called referencing
	changeMessage(&message)
}

func changeMessage(message *string) {
	// operation called de-referencing
	*message += "top"
	fmt.Println(*message) // output "Golang top"
}
