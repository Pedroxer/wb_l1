package main

import (
	"fmt"
)

// On устанавливает i-й бит в 1, используя побитовое ИЛИ
// с маской типа 00100 с единицей в позиции i
func On(n int64, i int) int64 {
	return n | 1<<(i-1)
}

// Off устанавливает i-й бит в 0, используя побитовое И
// с маской типа 11011 с нулем в позиции i
func Off(n int64, i int) int64 {
	return n & ^(1 << (i - 1))
}

func Demo() {

	var p int64 = 32

	on := On(p, 5)
	off := Off(p, 5)

	fmt.Printf("%b\n", p)
	fmt.Printf("%b\n", on)
	fmt.Printf("%b\n", off)

}
