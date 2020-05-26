package main

import (
	f "fmt"
	"sort"
)

func main() {

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	f.Println("Strings : ", strs)

	ints := []int{4, 7, 1}
	sort.Ints(ints)

	f.Println("Ints : ", ints)

	s := sort.IntsAreSorted(ints)
	f.Println("Sorted : ", s)

}
