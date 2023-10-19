package main

import (
	"fmt"
	"reflect"
)

func main() {
	var inter interface{} = make(chan int)
	switch inter.(type) {
	case int:
		fmt.Println("var is int")
		return
	case string:
		fmt.Println("var is string")
		return
	case bool:
		fmt.Println("var is bool")
		return
	case chan int:
	default:
		fmt.Println("unknown type. Maybe it's ", reflect.TypeOf(inter))
	}
}
