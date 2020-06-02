package main

import (
	f "fmt"
	"strings"
)

func main() {

	s := "Hello, Lucas"
	r := strings.NewReader(s)

	var s1, s2 string
	n, _ := f.Fscanf(r, "%s %s", &s1, &s2)

	f.Println("입력 개수 : ", n)
	f.Println(s1)
	f.Println(s2)
}
