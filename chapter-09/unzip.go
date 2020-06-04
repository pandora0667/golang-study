package main

import (
	"compress/gzip"
	f "fmt"
	"io/ioutil"
	"os"
)

func main() {

	file, err := os.Open("hello.txt.gz")
	if err != err {
		f.Println(err)
		return
	}
	defer file.Close()

	r, err := gzip.NewReader(file)
	if err != nil {
		f.Println(err)
		return
	}
	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		f.Println(err)
		return
	}

	f.Println(string(b))

}
