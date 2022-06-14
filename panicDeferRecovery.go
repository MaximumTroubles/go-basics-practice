package main

import "fmt"

func main() {
	// defer going to be called last in function, but can add additional execute times
	defer handePanic()
	Panic()
}

func Panic() {
	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}

	fmt.Println(len(messages), cap(messages))
	// here panic happened. When panic appears?
	// - when we go out of cap in array or slice
	// - when we are working with nil pointer
	// - when we cast incorrect types (interfaces)
	// - wrote down to close channel (?)
	// - or call it by ourselves panic("message")
	// exit status of panic eq 2

	//panic("help!")
	messages[4] = "message 5"
	fmt.Println(messages)
}

func Defer() {
	fmt.Println("I want to be called in the end of function just right before exiting")
}

func handePanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}

	fmt.Println("Panic: Panic handled successfully")
}
