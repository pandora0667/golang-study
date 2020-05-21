package main

import (
	f "fmt"
	"runtime"
	"sync"
)

func main() {

	f.Println("one time exec function")

	runtime.GOMAXPROCS(runtime.NumCPU())

	oneTime := new(sync.Once)

	for i := 0; i < 3; i++ {
		go func(n int) {
			f.Println("goroutine : ", n)
			oneTime.Do(Hello)
		}(i)
	}
	f.Scanln()
}

func Hello() {

	f.Println("Hello, World")

}
