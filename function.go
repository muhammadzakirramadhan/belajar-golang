package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func main() {
	for i := 0; i < 10; i++ {
		sayHello()
	}
}
