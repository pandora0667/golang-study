package main

import (
	"crypto/aes"
	f "fmt"
)

func main() {

	key := "lucas key 123456"
	s := "Hello, lucas! 12"

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		f.Println(err)
		return
	}

	cipherText := make([]byte, len(s))
	block.Encrypt(cipherText, []byte(s))
	f.Printf("%x\n", cipherText)

	planText := make([]byte, len(s))
	block.Decrypt(planText, cipherText)
	f.Printf(string(planText))
}
