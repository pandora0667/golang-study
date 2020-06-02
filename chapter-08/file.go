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
	w.WriteString("hello, wisoft")
	w.Flush()

	r := bufio.NewReader(file)
	fi, _ := file.Stat()
	b := make([]byte, fi.Size())

	file.Seek(0, os.SEEK_SET)
	r.Read(b)

	f.Println(string(b))
}
