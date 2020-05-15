package main

import f "fmt"

type rect struct {
	width, height int
}

func main() {

	f.Println("메서드")

	r := rect{width: 20, height: 5}
	f.Println("area : ", r.area())
	f.Println("perim : ", r.perim())

	rp := &r
	rp.height = 10
	rp.width = 5

	f.Println("area : ", rp.area())
	f.Println("perim : ", rp.perim())

}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}
