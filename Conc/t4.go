package Conc

import (
	"fmt"
	"github.com/Pedroxer/wbl_l1/util"
	"os"
	"os/signal"
	"syscall"
)

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
				data := util.RandomString(5)
				datach <- data
			}
		}
	}
}
