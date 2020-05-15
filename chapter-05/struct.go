package main

import f "fmt"

type person struct {
	name string
	age  int
}

func main() {

	f.Println("구조체")

	f.Println(person{"lucas", 27})
	f.Println(person{name: "bob"})
	f.Println(person{age: 19})

	s := person{name: "kan", age: 30}
	f.Println(s.name, s.age)

	f.Println(&s)

}
