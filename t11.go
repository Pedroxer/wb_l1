package main

import (
	"fmt"
	"math/rand"
)

func randMap() *map[int]struct{} {
	res := make(map[int]struct{})
	for i := 0; i < 5; i++ {
		res[rand.Intn(10)] = struct{}{}
	}
	return &res
}

func Intersect(map1, map2 *map[int]struct{}) map[int]struct{} {
	res := make(map[int]struct{})
	if len(*map1) > len(*map2) {
		map1, map2 = map2, map1
	}

	for i := range *map1 {
		if _, ok := (*map2)[i]; ok {
			res[i] = struct{}{}
		}
	}
	return res
}

func main() {
	map1 := randMap()
	map2 := randMap()
	fmt.Println("before intersect")
	fmt.Print("{ ")
	for key := range *map1 {
		fmt.Printf("%d ", key)
	}
	fmt.Println("}")

	fmt.Print("{ ")
	for key := range *map2 {
		fmt.Printf("%d ", key)
	}
	fmt.Println("}")

	map3 := Intersect(map1, map2)
	fmt.Print("After intersect { ")
	for key := range map3 {
		fmt.Printf("%d ", key)
	}
	fmt.Print("}")

}
