package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	zakir := Man{"Zakir"}
	zakir.Married()

	fmt.Println(zakir.Name)
}
