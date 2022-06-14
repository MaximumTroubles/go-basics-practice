package main

import (
	"fmt"
)

func main() {
	//array
	messages := [3]string{"1", "2", "3"}
	result := printMessages(messages)
	for _, i := range result {
		fmt.Printf(i + "\n")
	}
	fmt.Println(len(result))
}

func printMessages(messages [3]string) [3]string {
	messages[1] = "got it from print method"

	return messages
}
