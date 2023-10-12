package Conc

import (
	"fmt"
	"time"
)

func t5() {
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
			fmt.Println("Done") // Как только закончиться таймер, программа выйдет.
			return
		case res := <-dataCh:
			fmt.Println("From channel <-", res)
		}
	}

}
