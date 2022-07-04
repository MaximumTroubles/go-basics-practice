package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	go f(ch)
	time.Sleep(3 * time.Second)
	close(ch)
	time.Sleep(time.Second)
}

func f(stopCh chan struct{}) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-stopCh:
			return
		default:
		}

		select {
		case <-stopCh:
			return
		case <-ticker.C:
			fmt.Println("Task done")
		}
	}
}
