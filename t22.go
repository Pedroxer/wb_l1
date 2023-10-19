package main

import (
	"fmt"
	"math/big"
)

func sum(a, b *big.Int) *big.Int {
	var c big.Int
	c.Add(a, b)
	return &c
}

func sub(a, b *big.Int) *big.Int {
	var c big.Int
	c.Neg(b)
	return sum(a, &c)
}

func mul(a, b *big.Int) *big.Int {
	var c big.Int
	c.Mul(a, b)
	return &c
}

func div(a, b *big.Int) *big.Int {
	var c big.Int
	c.Quo(a, b)
	return &c
}
func main() {

	a := big.NewInt(1 << 30)
	b := big.NewInt(1 << 30 * 5)
	fmt.Println("sum = ", sum(a, b))
	fmt.Println("sub = ", sub(b, a))
	fmt.Println("mul = ", mul(a, b))
	fmt.Println("div = ", div(b, a))
}
