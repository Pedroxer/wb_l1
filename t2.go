package main

import (
	"fmt"
	"sync"
)

func calc(i int, mas []int, wg *sync.WaitGroup) {

	fmt.Printf("From goroutine %d send %d\n", i, mas[i]*mas[i])
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	mas := []int{2, 4, 6, 8, 10}
	for i := 0; i < len(mas); i++ {
		wg.Add(1)
		go calc(i, mas, &wg)
	}
	wg.Wait()
	fmt.Println("Done")
}
