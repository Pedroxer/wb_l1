package main

import (
	"fmt"

	"time"
)

func check(tm int, ch chan<- bool) {
	time.Sleep(time.Duration(tm) * time.Second)
	ch <- true
}

func main() {
	var (
		tm   int
		data string
	)
	fmt.Scan(&tm)

	timeout := time.After(time.Duration(tm) * time.Second)
	dataCh := make(chan string)

	go func() {
		for {
			fmt.Scan(&data)
			dataCh <- data
		}
	}()
	for {
		select {
		case <-timeout:
			fmt.Println("Done")
			return
		case res := <-dataCh:
			fmt.Println("From channel <-", res)
		}
	}

}
