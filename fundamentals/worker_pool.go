package main

import (
	"fmt"
	"time"
)

func workers(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func worker_pool_golang() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		fmt.Println("for loop w")
		go workers(w, jobs, results)
	}
	for j := 1; j <= numJobs; j++ {
		fmt.Println("for loop j")
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		fmt.Println("for loop a")
		fmt.Println("final result", <-results)
	}
}
