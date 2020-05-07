package main

import (
	f "fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	f.Println("현재 시간 : ", currentTime)

	weekday := time.Now().Weekday()

	switch weekday {
	case time.Monday:
		f.Println("월요일")
	case time.Tuesday:
		f.Println("화요일")
	case time.Thursday:
		f.Println("수요일")
	case time.Wednesday:
		f.Println("목요일")
	case time.Friday:
		f.Println("금요일")
	default:
		f.Println("주말")
	}
}
