package main

import (
	"compress/gzip"
	f "fmt"
	"os"
)

func main() {

	file, err := os.OpenFile(
		"hello.txt.gz",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))
	if err != nil {
		f.Println(err)
		return
	}
	defer file.Close()

	w := gzip.NewWriter(file)
	defer w.Close()

	s := "Hello World"
	w.Write([]byte(s))
	w.Flush()

}
