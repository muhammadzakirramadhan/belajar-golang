package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// pass by reference
	var address1 Address = Address{"Bekasi", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1
	// var address4 *Address = &Address{"Subang", "Jawa Barat", "Indonesia"}

	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	var alamat = Address{
		City:     "Makassar",
		Province: "Sulawesi Selatan",
		Country:  "",
	}

	ChangeCountryToIndonesia(&alamat)

	fmt.Println(alamat)
}
