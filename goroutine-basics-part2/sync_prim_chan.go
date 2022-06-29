package main

import (
	"fmt"
	"time"
)

func main() {
	// Not init channels equal nil
	message := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		message <-"Hello"
	}()

	msg := <-message
	fmt.Println(msg)
}
