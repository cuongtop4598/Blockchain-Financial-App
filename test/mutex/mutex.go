package main

import (
	"fmt"
	"sync"
	"time"
)

// mutex :  mutual exclusion
// SafeCounter is safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Look so only one goroutine at a time can access the map c.v
	defer c.mu.Unlock() // after the main function return a value, the child function will execute
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
