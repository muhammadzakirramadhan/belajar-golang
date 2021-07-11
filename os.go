package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	name, err := os.Hostname()

	if err == nil {
		fmt.Println("Hostname: ", name)
	} else {
		fmt.Println("Error Brow: ", err.Error())
	}

	fmt.Println("Arguments : ")
	fmt.Println(args)

	fmt.Println(args[1])
}
