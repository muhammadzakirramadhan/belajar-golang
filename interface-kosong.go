package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Upss"
	}
}

func main() {
	var data = Ups(3)

	fmt.Println(data)
}
