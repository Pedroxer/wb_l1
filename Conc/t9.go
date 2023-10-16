package Conc

import (
	"fmt"
	"sync"
)

func T9() {
	in := make(chan int)
	out := make(chan int)
	x := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() { // пишет из x в in
		for _, num := range x {
			in <- num
		}
		close(in)
		wg.Done()
	}()
	go func() { // пишет из in в out
		for i := range in {
			out <- i * 2
		}
		close(out)
		wg.Done()
	}()
	go func() { // вывод из out в os.Stdout
		for j := range out {
			fmt.Println(j)
		}
		wg.Done()
	}()
	wg.Wait()
}
