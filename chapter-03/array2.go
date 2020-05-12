package main

import f "fmt"

func main() {

	a := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		f.Println(a[i])
	}

	f.Println("--------------")

	for i, value := range a {
		f.Println(i, value)
	}

	f.Println("--------------")

	b := [...]string{"ios", "linux", "unix", "android", "windows"}
	for value := range b { // 인덱스 번호가 들어감
		f.Println(value)
	}

	f.Println("--------------")

	for _, value := range b {
		f.Println(value)
	}
}
