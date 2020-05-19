package main

import f "fmt"

func main() {

	f.Println("channels-range-close")

	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		f.Println(i)
	}

}
