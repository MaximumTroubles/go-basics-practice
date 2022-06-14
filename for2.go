package main

import "fmt"

func main() {
	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}

	// shor variant of cycle for. When we have slices, and we need to iterate each element it is a right way to do
	for index, message := range messages {
		fmt.Println(index, message)
	}
	// true cycle. for without condition
	counter := 0
	for {
		if counter == 100 {
			break
		}
		counter++
		fmt.Println(counter)
	}
}
