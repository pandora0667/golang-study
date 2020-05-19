package main

import (
	f "fmt"
	"runtime"
)

func main() {

	f.Println("goroutine multicore")

	runtime.GOMAXPROCS(runtime.NumCPU())
	f.Println(runtime.GOMAXPROCS(0))

	go LoopHello("Hello")
	f.Scanln()

}

func LoopHello(s string) {

	for i := 0; i < 100; i++ {
		f.Println(i+1, "번 : ", s)
	}

}
