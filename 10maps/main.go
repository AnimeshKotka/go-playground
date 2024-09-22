package main

import "fmt"

func main() {
	fmt.Println("Maps")
	langMap := make(map[string]string)

	langMap["JS"] = "Javascript"
	langMap["RB"] = "Ruby"
	langMap["PY"] = "Python"
	langMap["GO"] = "Golang"

	fmt.Println(langMap)

	// Delete from Map
	delete(langMap, "JS")

	langMap["TS"] = "TypeScript"
	fmt.Println(langMap)

	for key, value := range langMap {
		fmt.Printf("for key %v value is %v\n", key, value)
	}
}
