package main

import f "fmt"

func main() {

	f.Println("패닉 복구")

	printArray(1, 2, 3)
	f.Println("Hello, World")

}

func printArray(a int, b int, c int) {

	defer func() {
		s := recover()
		f.Println(s)
	}()

	array := [...]int{a, b, c}

	for i := 0; i < 5; i++ {
		f.Println(array[i])
	}

}
