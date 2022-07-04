package main

import "fmt"

const (
	openStatus  = "Open"
	closeStatus = "Close"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	close(ch)
	if value, open := <-ch; open {
		fmt.Printf("Number: %d - Status: %s \n", value, openStatus)
	}

	if value, open := <-ch; !open {
		fmt.Printf("Number: %d - Status: %s \n", value, closeStatus)
	}
}
