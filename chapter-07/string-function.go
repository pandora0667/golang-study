package main

import (
	f "fmt"
	"strings"
)

func main() {

	var p = f.Println

	p("Contains : ", strings.Contains("test", "es"))
	p("Count : ", strings.Count("test", "t"))
	p("HasPrefix : ", strings.HasPrefix("text", "te"))
	p("HasSuffix : ", strings.HasSuffix("test", "st"))
	p("Index : ", strings.Index("test", "t"))
	p("Index : ", strings.Index("test", "t"))
	p("Join : ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat : ", strings.Repeat("a", 5))
	p("Replace : ", strings.Replace("foo", "o", "0", -1))
	p("Replace : ", strings.Replace("foo", "o", "0", 1))
	p("Split : ", strings.Split("a-b-c-d-e", "-"))
	p("ToLower", strings.ToLower("TEST"))
	p("ToUpper : ", strings.ToUpper("test"))

	p("Len : ", len("hello"))
	p("Char : ", "hello"[1])

}
