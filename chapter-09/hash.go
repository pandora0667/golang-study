package main

import (
	"crypto/sha512"
	f "fmt"
)

func main() {

	s := "Hello, World"
	h1 := sha512.Sum512([]byte(s))
	f.Printf("%x\n", h1)

	sha := sha512.New()
	sha.Write([]byte("Hello, "))
	sha.Write([]byte("World"))

	h2 := sha.Sum(nil)
	f.Printf("%x\n", h2)

}
