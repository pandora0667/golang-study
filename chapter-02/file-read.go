package main

import (
	f "fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		f.Println("파일을 읽을 수 없습니다.")
		f.Println("Error Code : ", e)
	}
}

func CurrentDirectory() {
	path, _ := os.Getwd()
	f.Println(path)
}

func main() {

	CurrentDirectory()
	data, err := ioutil.ReadFile("./chapter-02/file/hello.txt")
	check(err)
	f.Print(string(data))
}
