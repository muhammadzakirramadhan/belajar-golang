package main

import "fmt"

func sayHello(firstname string, lastname string) {
	fmt.Println("Hallo ", firstname, lastname)
}

func main() {
	sayHello("Muhammad", "Zakir")
}
