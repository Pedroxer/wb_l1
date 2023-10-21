package main

import "fmt"

type Dog struct {
}

func (d *Dog) Wof() {
	fmt.Println("Wof, wof")
}

type Cat struct {
}

func (c *Cat) Meow(isCall bool) {
	if isCall {
		fmt.Println("Meow, Meow")
	}
}

// Dog и Cat это стороние библиотеки, которые мы не можем изменять. Цель - собрать слуховой адаптер, чтобы понять, кто, что говорит

type Adapter interface {
	Reaction()
}

type DogAdapter struct {
	*Dog
}

func NewDogAdapter(dog *Dog) Adapter {
	return &DogAdapter{dog}
}

func (d DogAdapter) Reaction() {
	d.Wof()
}

type CatAdapter struct {
	*Cat
}

func NewCatAdapter(cat *Cat) Adapter {
	return &CatAdapter{cat}
}

func (c CatAdapter) Reaction() {
	c.Meow(true)
}

// Так же есть жена, которая уже реализует интерфейс адаптера
type Wife struct{}

func (w Wife) Reaction() {
	fmt.Println("Привет, дорогой")
}

func main() {
	fmt.Println("\nВы останавливаетесь перед дверью и вставляете в ухо адаптер с двумя чипами\n")
	myFamily := [3]Adapter{NewDogAdapter(&Dog{}), NewCatAdapter(&Cat{}), Wife{}}

	fmt.Println("Открываете дверь и заходите домой\n")
	for _, member := range myFamily {
		member.Reaction()
	}
}
