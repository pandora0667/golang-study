package main

import (
	f "fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C // 타이머 만료되기 전까지 C 채널 블로킹
	f.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		f.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		f.Println("Timer 2 stopped")
	}
}
