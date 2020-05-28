package main

import (
	f "fmt"
	"reflect"
)

type Person struct {
	name string `tag1:"이름" tag2:"Name"`
	age  int    `tag1:"나이" tag2:"Age"`
}

func main() {

	p := Person{}

	name, ok := reflect.TypeOf(p).FieldByName("name")
	f.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2"))

	age, ok := reflect.TypeOf(p).FieldByName("age")
	f.Println(ok, age.Tag.Get("tag1"), age.Tag.Get("tag2"))

}
