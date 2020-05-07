package main

import f "fmt"

func main() {
	var input int
	f.Println("1~10 정수 입력")
	f.Scanln(&input)

	if input <= 10 {
		switch input {
		case 2, 4, 6, 8, 10:
			f.Println("짝수")
		default:
			f.Println("홀수")
		}
	} else {
		f.Println("잘못된 숫자 입력")
	}
}
