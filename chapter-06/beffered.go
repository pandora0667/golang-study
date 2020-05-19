package main

import f "fmt"

func main() {

	f.Println("beffered")

	messages := make(chan string, 2)
	messages <- "wisoft"
	messages <- "lab"

	f.Println(<-messages)
	f.Println(<-messages)

}
