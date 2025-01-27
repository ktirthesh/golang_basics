package main

import "fmt"

func channel_non_blocking_golang() {
	messages := make(chan string)
	signal := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "HI"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signal:
		fmt.Println("No signal received.", sig)
	default:
		fmt.Println("no activity")
	}
}
