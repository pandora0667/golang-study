package main

import (
	f "fmt"
	"runtime"
	"sync"
)

func main() {

	f.Println("wait group")

	runtime.GOMAXPROCS(runtime.NumCPU())

	Wait()
	f.Println("------------------------")
	DeferCall()

}

func Wait() {

	wait := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func(n int) {
			f.Println(n)
			wait.Done()
		}(i)
	}

	wait.Wait()
	f.Println("the end")

}

func DeferCall() {

	wait := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func(n int) {
			defer wait.Done()
			f.Println(n)
		}(i)
	}

	wait.Wait()
	f.Println("the end")

}
