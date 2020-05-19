package main

import f "fmt"

func main() {

	f.Println("goroutine 경량 실행 스레드")

	go Hello()
	f.Scanln()

}

func Hello() {

	f.Println("Hello World")

}
