package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "bCdefAaf"
	str = strings.ToLower(str)
	uniqueMap := make(map[byte]string)
	for i := 0; i < len(str); i++ {
		if _, ok := uniqueMap[str[i]]; ok {
			fmt.Println(false)
			return
		}
		uniqueMap[str[i]] = "1"

	}
	fmt.Println(true)

}
