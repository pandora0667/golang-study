package main

import f "fmt"

func main() {

	var a []int = make([]int, 5)
	var b = make([]string, 5)
	c := make([]float64, 5, 10)

	f.Println(a, b, c)

	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	f.Println(a)

	add := append(a, 6, 7, 8)
	f.Println("size :", len(add), ", components :", add)

	os := []string{"ios", "linux", "unix", "android", "windows"}
	f.Println(len(os), cap(os))
	f.Println(len(c), cap(c))

	f.Println(os[0:2])
	f.Println(os[1:3])
	f.Println(os[:3])
	f.Println(os[3:])

}
