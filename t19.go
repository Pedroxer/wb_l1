package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	runeStr := []rune(str)
	n := len(runeStr) - 1
	fmt.Println("Before ", str)
	for i := 0; i < n/2; i++ {

		runeStr[i], runeStr[n-i] = runeStr[n-i], runeStr[i]
	}
	fmt.Println("After ", string(runeStr))
}
