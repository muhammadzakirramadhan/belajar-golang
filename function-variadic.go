package main

import "fmt"

func sumAll(numbers ...int) int {
	t := 0

	for _, number := range numbers {
		t += number
	}

	return t
}

func main() {
	total := sumAll(10, 22, 13, 232, 23, 11)
	fmt.Println(total)

	slice := []int{21, 21, 33, 3, 13, 1, 31}
	total = sumAll(slice...)

	fmt.Println(total)
}
