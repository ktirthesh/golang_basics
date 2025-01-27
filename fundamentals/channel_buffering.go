package main

import "fmt"

func channel_buffering_golang() {
	message := make(chan string, 2)
	message <- "hello"
	message <- "World"
	fmt.Println(<-message)
	fmt.Println(<-message)
	// fmt.Println(<-message)

}
