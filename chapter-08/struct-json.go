package main

import (
	"encoding/json"
	f "fmt"
)

type Author struct {
	Name  string
	Email string
}

type Comments struct {
	Id      uint64
	Author  Author
	Content string
}

type Article struct {
	Id         uint64
	Title      string
	Author     Author
	Content    string
	Recommends []string
	Comments   []Comments
}

func main() {

	data := make([]Article, 1)

	data[0].Id = 1
	data[0].Title = "Hello, World"
	data[0].Author.Name = "lucas"
	data[0].Author.Email = "lucas@example.com"
	data[0].Content = "Good"
	data[0].Recommends = []string{"J", "K"}

	data[0].Comments = make([]Comments, 1)
	data[0].Comments[0].Id = 1
	data[0].Comments[0].Author.Name = "Keven"
	data[0].Comments[0].Author.Email = "keven@example.com"
	data[0].Comments[0].Content = "Nice keven"

	doc, _ := json.Marshal(data)
	f.Println(string(doc))

}
