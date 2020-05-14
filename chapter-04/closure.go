package main

import f "fmt"

func main() {

	f.Println("클로저")

	sum := func(a int, b int) int {
		return a + b
	}
	result := sum(1, 2)
	f.Println(result)

	f.Println("-----------------")

	cal := voltage()
	f.Println(cal(2, 4))
	f.Println(cal(5, 2))
	f.Println(cal(10, 5))

}

func voltage() func(i int, r int) int {

	return func(i int, r int) int {
		return i * r
	}

}
