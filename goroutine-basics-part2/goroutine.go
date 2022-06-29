package main

import (
	"fmt"
	"time"
)

func main() {
	timerChan := make(chan time.Time)
	go func() {
		time.Sleep(5 * time.Second)
		timerChan <- time.Now()
	}()

	fmt.Printf("%s\n", time.Now())

	completeAt := <- timerChan
	fmt.Println("from main routine " + time.Now().String()) 

	fmt.Printf("%s\n",completeAt)
}