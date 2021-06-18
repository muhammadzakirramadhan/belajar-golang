package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Zakir",
		"address": "Bekasi",
	}

	person["title"] = "Backend Developer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)

	book["title"] = "Belajar Golang"
	book["author"] = "Zakir"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
