package main

import (
	f "fmt"
	"time"
)

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go Worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {
		f.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		f.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
