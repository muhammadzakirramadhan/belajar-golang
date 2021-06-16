package main

import "fmt"

func main() {
	type NoKTP string
	type Gans bool

	var ktpZakir NoKTP = "123233333"
	var ganteng Gans = true

	fmt.Println(ktpZakir)
	fmt.Println(NoKTP("132323444"))
	fmt.Println("Zakir Gans? ", ganteng)
}
