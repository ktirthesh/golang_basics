package main

import "fmt"

func channels_golang() {
	message := make(chan string)
	go func() { message <- "ping" }()
	msg := <-message
	fmt.Println(msg)
}
