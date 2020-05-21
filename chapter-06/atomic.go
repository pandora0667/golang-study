package main

import (
	f "fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {

	f.Println("atomic")

	runtime.GOMAXPROCS(runtime.NumCPU())

	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for true {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	f.Println("ops : ", opsFinal)
}
