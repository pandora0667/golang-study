package main

import (
	"errors"
	f "fmt"
	"math"
)

func main() {

	f.Println("err")
	result, err := Sqrt(0)

	f.Println(result)
	f.Println(err)

}

func Sqrt(f float64) (float64, error) {

	if f <= 0 {
		return 0, errors.New("math : source root of negative number")
	} else {
		return math.Sqrt(f), errors.New("ok")
	}

}
