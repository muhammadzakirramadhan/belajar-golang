package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("Z([a-z])k")

	fmt.Println(regex.MatchString("Zak"))
	fmt.Println(regex.MatchString("Zik"))
	fmt.Println(regex.MatchString("ZAK"))

	fmt.Println(regex.FindAllString("Zak Zik Zok ZAk", -1))
}
