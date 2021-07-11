package main

import (
	"belajar-golang/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
