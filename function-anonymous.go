package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Your Are Blacklist", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// func blacklistZakir(name string) bool {
// 	return name == "Zakir"
// }

func main() {

	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Zakir", blacklist)
	registerUser("anjing", blacklist)

	registerUser("admin", func(name string) bool {
		return name == "admin"
	})

}
