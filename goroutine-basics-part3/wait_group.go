package main

import (
	"fmt"
	"sync"
)

const coutner int = 1000

func main() {
	// wait group it is a stucture
	wg := sync.WaitGroup{}

	v := 0
	for i := 0; i < coutner; i++ {
		wg.Add(1)
		go func() {
			v++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(v)
}
