package main

import "fmt"

func goodbye(name string) string {
	return "Good Bye " + name
}

func main() {
	sayGoodBye := goodbye
	result := sayGoodBye("Zakir")

	fmt.Println(result)

}
