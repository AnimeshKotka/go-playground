package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditional if else cases")

	var email string = "example.com"

	length := len(email)
	if length > 0 {
		fmt.Println("email is valid")
	} else {
		fmt.Println("Invalid")
	}

	// With shorter form
	if len(email) > 0 {
		fmt.Println("email is valid")
	} else {
		fmt.Println("Invalid")
	}
}
