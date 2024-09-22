package main

import "fmt"

func main() {

	// Var style where we use var key before var name and then type

	// string variables

	var name string = "mjjdifj"

	fmt.Println(name)
	fmt.Printf("this is type pf %T  \n", name)

	// boolena variables

	var isOld bool = false

	fmt.Println(isOld)
	fmt.Printf("this is type pf %T  \n", isOld)

	// number variables

	var age int = 35

	fmt.Println(age)
	fmt.Printf("this is type pf %T  \n", age)

	// float variables

	var height float32 = 6.7444444444444444444444444444
	var weight float64 = 6.7444444444444444444444444444

	fmt.Println(height, weight)
	fmt.Printf("this is type pf %T %T \n", height, weight)

	// alias variables

	var alisa rune = 23

	fmt.Println(alisa)
	fmt.Printf("this is type pf %T  \n", alisa)

	// No Var style
	// 1. is only used inside a method, global variable not allowed
	// 2. go lexer will change the type accordingly
	class := 12
	fmt.Println(class)
}
