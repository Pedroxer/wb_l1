package Conc

import (
	"fmt"
	"sync"
)

func calc2(i int, mas []int, wg *sync.WaitGroup, out chan int) {

	fmt.Printf("From goroutine %d send %d\n", i, mas[i]*mas[i])
	out <- (mas[i] * mas[i])
	wg.Done()
}

func Conc3() {
	var wg sync.WaitGroup
	mas := []int{2, 4, 6, 8, 10}
	sum := 0
	ch := make(chan int)
	for i := 0; i < len(mas); i++ {
		wg.Add(1)
		go calc2(i, mas, &wg, ch)
		sum += <-ch
	}
	wg.Wait()
	fmt.Println("Sum = ", sum)
}
