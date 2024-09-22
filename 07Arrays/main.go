package main

import "fmt"

func main() {

	// Arrays

	var arr = []string{} // declare an with must have len

	// arr[0] = "Hello"
	// arr[1] = "World"

	fmt.Println(arr)

	arr = append(arr, "test")
	fmt.Println(arr)

	// declare and initialize

	arr2 := [2]string{"Hello", "World"}
	fmt.Println(arr2)

	// declare and initialize with len

	arr3 := [...]string{"Hello", "World"}
	fmt.Println(arr3)

	// slice

	slice := []string{"Hello", "World"}
	fmt.Println(slice)

	slice = append(slice, "Go") // append: it reallocates the default memory and create a new memo block
	// with updated values
	fmt.Println(slice)

	// slice operations

	fmt.Println(slice[1:2])

	// remove value from slice
	slice = append(slice[:2], slice[3:]...) // remove item from slice
	fmt.Println(slice)

}
