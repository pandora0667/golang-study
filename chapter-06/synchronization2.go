package main

import (
	f "fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	Condition() // 경쟁상황으로 인해 제대로 된 값이 나오지 않음
	Mutex()
}

func Condition() {

	var data = []int{}
	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1)
			runtime.Gosched()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1)
			runtime.Gosched()
		}
	}()

	time.Sleep(2 + time.Second)
	f.Println(len(data))

}

func Mutex() {

	var data = []int{}
	var mutex = new(sync.Mutex)

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()

			runtime.Gosched()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()

			runtime.Gosched()
		}
	}()

	time.Sleep(2 * time.Second)
	f.Println(len(data))

}
