package main

import (
	"bufio"
	f "fmt"
	"os"
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

	w := bufio.NewWriter(file)
	f.Fprintf(w, "%d,%f,%s", 1, 1.1, "Hello")
	w.Flush()
}
