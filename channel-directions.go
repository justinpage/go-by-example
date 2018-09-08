package main

import "fmt"

// This ping function only accepts a channel for sending values. It would be a
// compile-time error to try to receive on this channel.
func Ping(pings chan<- string, msg string) {
	pings <- msg
}

// The pong function accepts one channel for receives (pings) and a second for
// sends (pongs).
func Pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	Ping(pings, "passed message")
	Pong(pings, pongs)
	fmt.Println(<-pongs)
}
