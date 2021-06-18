package main

import "fmt"

func getFullname() (string, string) {
	return "Muhammad", "Zakir"
}

func main() {
	firstname, lastname := getFullname()

	fmt.Println(firstname)
	fmt.Println(lastname)

}
