package main

import (
	"container/list"
	f "fmt"
)

func main() {

	l := list.New()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	f.Println("Front : ", l.Front().Value)
	f.Println("Back : ", l.Back().Value)

	for e := l.Front(); e != nil; e = e.Next() {
		f.Println(e.Value)
	}
}
