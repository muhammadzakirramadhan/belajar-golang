package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Zakir Wibu", "Wibu"))
	fmt.Println(strings.Split("Zakir Wibu", " "))
	fmt.Println(strings.ToLower("Zakir WIBU"))
	fmt.Println(strings.ToUpper("Zakir WIBU"))
	fmt.Println(strings.ToTitle("Zakir WIBU"))
	fmt.Println(strings.Trim("	Zakir WIBU	", "	"))
	fmt.Println(strings.ReplaceAll("Zakir Wibu", "Wibu", "Ganteng"))
}
