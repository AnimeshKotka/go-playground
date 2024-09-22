package main

import "fmt"

func main() {
	// Loops

	// for
	// for init; condition; post {
	// }
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println("For Loop")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// for range
	fmt.Println("For Range")
	for index, day := range days {
		fmt.Println(index, day)
	}

	// while
	fmt.Println("While Loop")
	i := 0
	for i < len(days) {
		fmt.Println(days[i])
		i++
	}
}
