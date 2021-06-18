package main

import "fmt"

func hay(name string) string {
	if name == "" {
		return "Halo Bro"
	} else {
		return "Halo " + name
	}
}

func main() {
	result := hay("Zakir")

	fmt.Println(result)
}
