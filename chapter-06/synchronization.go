package main

import (
	f "fmt"
	"time"
)

func main() {

	f.Println("synchronization")

	done := make(chan bool, 1)
	go worker(done)

	<-done
}

func worker(done chan bool) {

	f.Println("working...")
	time.Sleep(time.Second)
	f.Println("done")

	done <- true
}
