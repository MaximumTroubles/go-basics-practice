package main

import (
	"fmt"
	"sync"
	"time"
)

type SaveCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SaveCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SaveCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SaveCounter{v: make(map[string]int)}
	for i := 0; i < 100000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
