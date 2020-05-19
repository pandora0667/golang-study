package main

import f "fmt"

func main() {

	msg := make(chan string)

	go func() { msg <- "ping" }()

	result := <-msg
	f.Println(result)

}
