package main

import "fmt"

func main() {
	in := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	res := make(map[int][]float64)
	for i := range in {
		buk := int((int(in[i]) / 10) * 10) // узнаем ключ с шагом 10
		res[buk] = append(res[buk], in[i])
	}
	for key, value := range res {
		fmt.Printf("%d:{%v}\n", key, value)
	}
}
