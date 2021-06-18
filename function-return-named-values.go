package main

import "fmt"

func getFullName() (firstname, middlename, lastname string) {
	firstname = "Muhammad"
	middlename = "Zakir"
	lastname = "Ramadhan"

	return
}

func main() {
	firstname, middlename, lastname := getFullName()

	fmt.Println(firstname)
	fmt.Println(middlename)
	fmt.Println(lastname)

}
