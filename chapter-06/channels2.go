package main

import f "fmt"

func main() {

	f.Println("Channels")

	c := make(chan int)
	go sum(10, 20, c)

	result := <-c

	f.Println(result)

}

func sum(a int, b int, c chan int) {

	c <- a + b

}
