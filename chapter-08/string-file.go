package main

import (
	"bufio"
	f "fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))

	if err != nil {
		f.Println(err)
		return
	}
	defer file.Close()

	s := "Hello, World"
	r := strings.NewReader(s)

	w := bufio.NewWriter(file)

	w.ReadFrom(r)
	w.Flush()

}
