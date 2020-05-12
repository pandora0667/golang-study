package main

import f "fmt"

func main() {

	original := [...]string{"ios", "linux", "unix", "android", "windows"}
	copy := original // 배열 전체 복사

	f.Println(original)
	f.Println(copy)
}
