package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter age")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	numAge, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	fmt.Printf("age type is %T and value is %v\n", input, input)
	fmt.Printf("age type is %T and value is %v\n", numAge, numAge)

}
