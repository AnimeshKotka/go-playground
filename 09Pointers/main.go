package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Pointer\n")

	var ptr *int // If we don't initialize the pointer it will be nil
	fmt.Printf("ptr: %v\n", ptr)

	number := 10
	ptr = &number

	fmt.Printf("ptr of a actual value: %v\n", ptr)
	fmt.Printf("ptr value is: %v\n", *ptr)
}
