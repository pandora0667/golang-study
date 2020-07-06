package main

import (
	f "fmt"
	"runtime"
)

func main() {

	f.Println("goroutun-closures")

	runtime.GOMAXPROCS(1) // single core

	s := "wisoft"

	for i := 0; i < 100; i++ {
		go func(number int) {
			f.Println(s, number)
		}(i)
	}

	f.Scanln()
}
