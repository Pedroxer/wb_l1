package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	s := strings.Fields(str)

	n := len(s) - 1
	fmt.Println("Before ", str)
	for i := 0; i < n/2; i++ {

		s[i], s[n-i] = s[n-i], s[i]
	}
	fmt.Println("After ", strings.Join(s, " "))
}
