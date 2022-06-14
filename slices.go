package main

import "fmt"

func main() {
	messages := make([]int, 5, 99)
	messages = append(messages, 1)
	fmt.Println(messages)
	fmt.Println(len(messages))
	fmt.Println(cap(messages))
}
