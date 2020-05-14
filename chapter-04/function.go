package main

import f "fmt"

func main() {

	a := 10
	b := 20

	result := sum(a, b)
	f.Println(result)

	f.Println(subtract(30, 20))

	f.Println(division(4, 3))
	_, remainder := division(456, 25)
	f.Println("나머지 : ", remainder)

}

func sum(a int, b int) int {
	return a + b
}

func subtract(a int, b int) (r int) {
	r = a - b
	return r
}

func division(a int, b int) (int, int) {
	value := a / b
	remainder := a % b

	return value, remainder
}
