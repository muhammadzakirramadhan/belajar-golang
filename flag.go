package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your host")
	var user *string = flag.String("user", "root", "Put your username")
	var password *string = flag.String("password", "toor", "Put your password")

	flag.Parse()

	fmt.Println("Host: ", *host)
	fmt.Println("User: ", *user)
	fmt.Println("Password: ", *password)
}
