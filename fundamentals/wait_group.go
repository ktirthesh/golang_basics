package main

import (
	"fmt"
	"sync"
	"time"
)

func work(id int) {
	fmt.Printf("Worker %d string \n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func waitgroup_golang() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work(i)
		}()
	}
	wg.Wait()
}
