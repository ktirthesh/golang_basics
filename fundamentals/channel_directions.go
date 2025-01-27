package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	fmt.Println("pong msg", msg)
	pongs <- msg
}
func channel_directions_golang() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "string passed")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
