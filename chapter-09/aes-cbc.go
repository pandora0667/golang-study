package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	f "fmt"
	"io"
)

func main() {

	key := "Hello Key 123456"

	s := `동해 물과 백두산이 마르고 닳도록
하느님이 보우하사 우리나라 만세.
무궁화 삼천리 화려강산
대한 사람, 대한으로 길이 보전하세.`

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		f.Println(err)
		return
	}

	cipherText := encrypt(block, []byte(s))
	f.Printf("%x\n", cipherText)

	plainText := decrypt(block, cipherText)
	f.Printf(string(plainText))
}

func encrypt(block cipher.Block, plainText []byte) []byte {

	if mod := len(plainText) % aes.BlockSize; mod != 0 {
		padding := make([]byte, aes.BlockSize-mod)
		plainText = append(plainText, padding...)
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		f.Println(err)
		return nil
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return cipherText

}

func decrypt(block cipher.Block, cipherText []byte) []byte {

	if len(cipherText)%aes.BlockSize != 0 {
		f.Printf("암호화된 데이터의 길이는 블록 크기의 배수가 되어야 합니다.")
		return nil
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	plainText := make([]byte, len(cipherText))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plainText, cipherText)

	return plainText

}
