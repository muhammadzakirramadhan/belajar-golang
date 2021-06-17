package main

import "fmt"

func main() {
	var ujian = 90
	var absensi = 70

	var lulusUjian = ujian > 80
	var lulusAbsen = absensi > 60

	var lulus bool = lulusUjian && lulusAbsen

	fmt.Println(lulus)

	fmt.Println(ujian > 80 && absensi > 60)
}
