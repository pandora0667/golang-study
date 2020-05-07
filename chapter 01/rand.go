package main

import (
	f "fmt"
	"math/rand"
	_ "net/http"
	"time"
)

func main() {
	a := 10
	b := 20
	_ = b

	f.Println(a)

	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	f.Println(random.Intn(100))
}
