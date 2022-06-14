package main

import (
	"fmt"
	"sync"
)

// mutex Мьютексы позволяют разграничить доступ к некоторым общим ресурсам, гарантируя,
// что только одна горутина имеет к ним доступ в определенный момент времени.

// common res.
var counter int = 0

func main() {

	ch := make(chan bool)
	var mutex sync.Mutex
	for i := 1; i < 5; i++ {
		go work(i, ch, &mutex)
	}

	// waiting for all routines end
	for i := 1; i < 5; i++ {
		<-ch
		fmt.Println("what happen here", i)
	}

	fmt.Println("The End")
}

func work(n int, ch chan bool, mutex *sync.Mutex) {
	// We locked mutex here so it mean we close access to further code for others  routines until it unlock
	// First routine that get here lock it and use other in queue, second after it unlock can reach code, rest routice again in queue
	mutex.Lock() 
	counter = 0
	for i := 1; i <= 5; i++ {
		counter++
		fmt.Println("Goroutine", n, " - ", counter )
	}
	mutex.Unlock()
	ch <- true
}