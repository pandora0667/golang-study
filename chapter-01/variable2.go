package main

import f "fmt"

type ByteSize uint64

var (
	i int
	b bool
	s string
)

var name, title, num1, num2 = "seongwon", "golang", 1, 2

const NICKNAME = "lucas"

const (
	GO = iota // 여러 상수를 열거하고, 0부터 1씩 값을 증가시킨다.
	JAVA
	PYTHON
	C
)

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PT
	EB
)

func main() {
	i, b, s = 1, true, "example"

	f.Println(i, b, s)
	f.Println("nickname : ", NICKNAME)
	f.Println(name, title, num1, num2)

	f.Println(GO, JAVA, PYTHON, C)

	f.Println(KB, MB, GB, TB, PT, EB)
}
