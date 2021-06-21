package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name Is", customer.Name)
}

func (cs Customer) sayHu() {
	fmt.Println("Huuuu from", cs.Name)
}

func main() {
	var zakir Customer

	zakir.Name = "Zakir"
	zakir.Address = "Bekasi"
	zakir.Age = 21

	zakir.sayHi("zakir")
	zakir.sayHu()

	fmt.Println(zakir.Name)
	fmt.Println(zakir.Address)
	fmt.Println(zakir.Age)

	yuzi := Customer{
		Name:    "Yuzi",
		Address: "Bekasi",
		Age:     20,
	}

	fmt.Println(yuzi)

	faza := Customer{"Faza", "Bekasi", 20}

	fmt.Println(faza)
}
