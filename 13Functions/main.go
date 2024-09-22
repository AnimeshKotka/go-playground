package main

import "fmt"

func main() { // this is main function, entry point of the program
	fmt.Println("Functions")

	addRes := addTwoNumber(10, 20)
	fmt.Println("after addition", addRes)

	multiAddRes := addMultiple(2, 3, 45, 6, 7, 2)
	fmt.Println("after multiple addition", multiAddRes)

	a := 10
	b := 10

	pointerAddRes := addUsingPointer(&a, &b)
	fmt.Println("after pointer addition", pointerAddRes)
}

func addTwoNumber(a int, b int) int {
	return a + b
}

func addMultiple(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func addUsingPointer(a *int, b *int) int {
	return *a + *b
}
