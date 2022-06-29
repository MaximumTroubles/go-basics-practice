package main

import (
	"fmt"
	"time"
)

func main() {
	message1 := make(chan string)
	message2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			message1 <- "half second have passed"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			message2 <- "2 second have passed"
		}
	}()

	for {
		select {
		case msg := <-message1:
			fmt.Println(msg)

		case msg := <-message2:
			fmt.Println(msg)
		}
	}
}
