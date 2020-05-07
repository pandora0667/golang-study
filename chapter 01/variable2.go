package main

import f "fmt"

var (
	i int
	b bool
	s string
)

const nickname = "lucas"

var name, title, num1, num2 = "seongwon", "golang", 1, 2

const (
	golang = iota
	java
	python
	c
)

func main() {
	i, b, s = 1, true, "example"

	f.Println(i, b, s)
	f.Println("nickname : ", nickname)
	f.Println(name, title, num1, num2)

	f.Println(golang, java, python, c)
}
