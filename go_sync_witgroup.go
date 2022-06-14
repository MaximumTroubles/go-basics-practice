package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	work := func (id int)  {
		defer wg.Done()
		fmt.Printf("Goroutin %d begin to perform \n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Goroutin %d end up to perform \n", id)
	}

	go work(1)
	go work(2)

	wg.Wait()
	fmt.Println("Goroutines end up to perform")
}