package main

import (
	f "fmt"
	"time"
)

func main() {

	f.Println("select")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			f.Println("received", msg1)
		case msg2 := <-c2:
			f.Println("received", msg2)
		}
	}

}
