package main

import (
	f "fmt"
	"sort"
)

type ByLength []string

func main() {

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	f.Println(fruits)

}

func (s ByLength) Len() int {

	return len(s)
}

func (s ByLength) Swap(i, j int) {

	s[i], s[j] = s[j], s[i]

}

func (s ByLength) Less(i, j int) bool {

	return len(s[i]) < len(s[j])

}
