package main

import (
	"container/ring"
	f "fmt"
)

func main() {

	data := []string{"lucas", "john", "maria", "James"}

	r := ring.New(len(data))
	for i := 0; i < r.Len(); i++ {
		r.Value = data[i]
		r = r.Next()
	}

	r.Do(func(x interface{}) { // 각 요소에 대해 실행할 함수를 지정해 두면 링을 구성하는 모든 데이터 각각에 대해 실행할 함수를 호출
		f.Println(x)
	})

	f.Println("Move Forward : ")
	r = r.Move(1)

	f.Println("Curr : ", r.Value)
	f.Println("Prev : ", r.Prev().Value)
	f.Println("Next : ", r.Next().Value)
}
