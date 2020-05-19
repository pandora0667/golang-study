package main

import (
	f "fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	Condition()
	f.Println("-------------------")
	Mutex()

}

func Condition() {

	data := 0

	go func() {
		for i := 0; i < 3; i++ {
			data += 1
			f.Println("write : ", data)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			f.Println("read : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			f.Println("read2 : ", data)
			time.Sleep(2 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)

}

func Mutex() {

	data := 0
	rwMutex := new(sync.RWMutex)

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock()
			data += 1
			f.Println("write : ", data)
			time.Sleep(10 * time.Millisecond)
			rwMutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock()
			f.Println("read : ", data)
			time.Sleep(1 * time.Second)
			rwMutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock()
			f.Println("read2 : ", data)
			time.Sleep(2 * time.Second)
			rwMutex.Unlock()
		}
	}()

	time.Sleep(10 * time.Second)

}
