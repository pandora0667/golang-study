package main

import f "fmt"

func main() {

	var empty [5]int
	empty[4] = 5
	f.Println(empty)

	var a [3]int = [3]int{1, 2, 3}
	var b = [3]string{"name", "age", "weight"}
	var c = [5]float32{
		1.14,
		2.14,
		3.14,
		4.14,
		5.14,
	}

	d := [...]string{"apple", "lg", "samsung", "kt"}

	f.Println("array size", len(a))

	for i := 0; i < len(b); i++ {
		f.Println("b 내용 : ", b[i])
	}

	f.Println(c, d)

	var multiArray = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	f.Println("----------------------")
	f.Println(multiArray)
	f.Println("array size", len(multiArray))

}
