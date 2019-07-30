package main

import "fmt"

// Accepts a channel for sending values, not receiving them
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Accepts a channel for receiving values, and another for sending values
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	go ping(pings, "passed message")
	go pong(pings, pongs)

	fmt.Println(<-pongs)
}
