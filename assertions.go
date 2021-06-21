package main

import (
	"fmt"
)

func random() interface{} {
	return "OK"
}

func main() {
	hasil := random()
	// resultString := hasil.(string)
	// fmt.Println(resultString)

	switch value := hasil.(type) {
	case int:
		fmt.Println("Int Value", value)
	case string:
		fmt.Println("String Value", value)
	default:
		fmt.Println("Unknown Types")
	}
}
