package main

import (
	"encoding/json"
	f "fmt"
)

func main() {

	doc := `{"name": "lucas", "age": "27"}`

	var data map[string]interface{}

	json.Unmarshal([]byte(doc), &data)
	f.Println(data["name"], data["age"])

	f.Println("------------------")
	json2()

}

func json2() {

	data := make(map[string]interface{})

	data["name"] = "K"
	data["age"] = 10

	doc, _ := json.Marshal(data)
	f.Println(string(doc))
}
