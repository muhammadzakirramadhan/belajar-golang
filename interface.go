package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Anime struct {
	Name string
}

func sayHelo(hn HasName) {
	fmt.Println("Hello", hn.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (a Anime) GetName() string {
	return a.Name
}

func main() {
	var zakir Person
	zakir.Name = "Zakir"

	sayHelo(zakir)

	bofuri := Anime{
		Name: "Bofuri",
	}

	sayHelo(bofuri)
}
