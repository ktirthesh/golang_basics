package main

import (
	"fmt"
	"time"
)

func fff(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func coroutine_golang() {
	fff("direct")

	go fff("goroutinefdfd")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
