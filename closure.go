package main

import "fmt"

func main() {
	counter := 0
	name := "Zakir"

	increment := func() {
		name := "Yuzi"
		fmt.Println("Increment")
		fmt.Println(name)
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
