package main

import (
	"fmt"
	"math/rand"

	"github.com/yourbasic/graph"
)

func main() {
	n := rand.Intn(100)
	g := graph.New(n)
	fmt.Println(g)
}
