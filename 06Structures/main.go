package main

import "fmt"

type message struct {
	phone   int
	message string
}

func main() {

	var msg1 message = message{
		12345678,
		"nbjbdejsnxjkswhk",
	}

	fmt.Println(msg1.message, msg1.phone)
}
