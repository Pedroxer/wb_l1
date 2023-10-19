package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	C int
	m sync.Mutex
}

func main() {
	c := &Counter{C: 0}

	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		go func() {
			c.m.Lock()
			c.C++
			c.m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.C)
}
