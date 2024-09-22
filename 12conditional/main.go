package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Conditional statements")

	isLogout := true

	var result string

	if isLogout {
		result = "User is logged out"
	} else {
		result = "User is logged in"
	}

	fmt.Println(result)

	rand.Seed(time.Now().UnixNano())
	if num := rand.Intn(999); num%2 == 0 {
		fmt.Println("Even number", num)
	} else {
		fmt.Println("Odd number", num)
	}
}
