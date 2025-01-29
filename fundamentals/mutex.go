package main

import (
	"fmt"
	"sync"
)

type Container1 struct {
	mu      sync.Mutex
	counter map[string]int
}

func (c *Container1) inc(name string) {
	c.mu.Lock()
	fmt.Printf("locked and counter added %s \n", name)
	defer c.mu.Unlock()
	c.counter[name]++
}

func mutex_golang() {
	c := Container1{
		counter: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup
	doIncremental := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	wg.Add(3)
	go doIncremental("a", 10000)
	go doIncremental("a", 10000)
	go doIncremental("b", 10000)
	wg.Wait()
	fmt.Println(c.counter)

}
