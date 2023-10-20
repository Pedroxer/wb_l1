package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d) // Выполнение блокируется, пока After не вернёт данные в канал.
}

func main() {
	fmt.Println("Before")
	sleep(10 * time.Second)
	fmt.Println("After")
}
