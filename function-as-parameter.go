package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	namedFilter := filter(name)

	fmt.Println("Hello ", namedFilter)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Zakir", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}
