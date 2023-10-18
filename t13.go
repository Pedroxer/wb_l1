package main

import "fmt"

func Swap() {
	a := 5
	b := 6
	fmt.Println("Before a = ", a, " ", "b = ", b)

	a, b = b, a
	fmt.Println("After v1 a = ", a, " ", "b = ", b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Println("After v2 a = ", a, " ", "b = ", b)
}
