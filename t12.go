package main

import "fmt"

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})
	for _, val := range str {
		set[val] = struct{}{}
	}

	for key := range set {
		fmt.Print(key, " ")
	}
}
