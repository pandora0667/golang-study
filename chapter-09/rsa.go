package main

import (
	"crypto/rand"
	"crypto/rsa"
	f "fmt"
)

func main() {

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		f.Println(err)
		return
	}

	publicKey := &privateKey.PublicKey

	s := `동해 물과 백두산이 마르고 닳도록
하느님이 보우하사 우리나라 만세.
무궁화 삼천리 화려강산
대한 사람, 대한으로 길이 보전하세.`

	cipherText, err := rsa.EncryptPKCS1v15(
		rand.Reader,
		publicKey,
		[]byte(s))

	f.Printf("%x\n", cipherText)

	plainText, _ := rsa.DecryptPKCS1v15(
		rand.Reader,
		privateKey,
		cipherText,
	)

	f.Println(string(plainText))

}
