package main

import f "fmt"

func main() {

	f.Println("Channel direction")

	pings := make(chan string, 2)
	pongs := make(chan string, 2)

	ping(pings, "passed message")
	pong(pings, pongs)

	f.Println(<-pongs)

}

func ping(pings chan<- string, msg string) {

	pings <- msg

}

func pong(pings <-chan string, pongs chan<- string) {

	msg := <-pings
	pongs <- msg

}
