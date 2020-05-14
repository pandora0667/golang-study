package main

import f "fmt"

func main() {

	f.Println("포인터")

	var numPtr *int
	f.Println(numPtr)
	var numPtr2 *int = new(int)
	f.Println(numPtr2)

	i := 1
	f.Println("init : ", i)
	number(i)
	f.Println("value", i)
	pointerNumber(&i)
	f.Println("pointer : ", i)
	f.Println("address : ", &i)
}

func number(a int) {
	a = 0
}

func pointerNumber(a *int) {
	*a = 0
}
