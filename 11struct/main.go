package main

import "fmt"

func main() {
	fmt.Println("Structs")

	p1 := Person{"James", "Bond", 32}
	p2 := Person{"Miss", "Moneypenny", 27}

	fmt.Printf("%+v %v", p1, p2)
	fmt.Println("ages", p1.Age, p2.Age)
}

type Person struct {
	First string
	Last  string
	Age   int
}
