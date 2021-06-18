package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Loop ke ", counter)
	}

	slice := []string{"zakir", "azzah", "zak"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for _, value := range slice {
		// fmt.Println("Index ", index, " = ", value)
		fmt.Println(value)
	}

	person := make(map[string]string)

	person["name"] = "Zakir"
	person["title"] = "Backend Developer"

	for key, val := range person {
		fmt.Println(key, " = ", val)
	}
}
