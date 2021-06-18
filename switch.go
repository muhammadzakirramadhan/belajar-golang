package main

import "fmt"

func main() {
	name := "Zakir"

	switch name {
	case "Zakir":
		fmt.Println("Bener")

	case "Ramadhan":
		fmt.Println("Gak Dlu")

	default:
		fmt.Println("Hadeh")
	}

	// switch length := len(name); length > 4 {
	// case true:
	// 	fmt.Println("Panjangan Woi")
	// case false:
	// 	fmt.Println("Kependekan Woi")
	// }

	length := len(name)

	switch {
	case length > 5:
		fmt.Println("Kepanjangan Woi")
	case length < 5:
		fmt.Println("Kependekan Woi")
	default:
		fmt.Println("Bener Pas")
	}

}
