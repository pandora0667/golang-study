package main

import (
	f "fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			f.Println("Tick at ", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	f.Println("Ticker stopped")
}
