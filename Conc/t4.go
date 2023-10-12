package Conc

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func randomString(n int) string {
	var key = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	res := ""
	for i := 0; i < n; i++ {
		res += string(key[rand.Intn(len(key))])
	}
	return res
}

func Conc4() {
	fmt.Print("N = ")
	var worker int
	fmt.Scan(&worker)
	inter := make(chan os.Signal, 1)
	signal.Notify(inter, os.Interrupt, syscall.SIGTERM)
	for true {

		datach := make(chan string, worker)

		for w := 1; w <= worker; w++ {
			go func(i int, ch chan string) {
				for j := range ch {
					fmt.Printf("Worker %d:%s\n", i, j)

				}
			}(w, datach)
		}
		for { // Дожидаюсь, пока из буфера всё прочту, чтобы не потерять информацю в канале.
			select {
			case <-inter:
				for len(datach) != 0 {
				}
				fmt.Println()
				return

			default:
				data := randomString(5)
				datach <- data
			}
		}
	}
}
