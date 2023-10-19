package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	i := 2
	a = append(a[:i], a[i+1:]...)
	fmt.Println(a)
}
