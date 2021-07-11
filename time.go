package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Nanosecond())
	fmt.Println(now.Second())
	fmt.Println(now.Minute())
	fmt.Println(now.Hour())
	fmt.Println(now.Day())
	fmt.Println(now.Month())
	fmt.Println(now.Year())

	utc := time.Date(2021, 06, 06, 06, 06, 06, 06, time.UTC)

	fmt.Println(utc)

	layout := "2006-06-01"
	parse, _ := time.Parse(layout, "2021-06-06")

	fmt.Println(parse)
}
