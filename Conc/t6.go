package Conc

import (
	"context"
	"fmt"
	"sync"
)

func Goroutins() {
	var wg *sync.WaitGroup
	// Читающая из канала горутина завершит работу при закрытии канала
	wg.Add(1)
	op := make(chan int)
	go func() {
		fmt.Println("routine start")
		for i := range op {
			fmt.Println("read", i)
		}
		fmt.Println("routine stops")
		wg.Done()
	}()
	for i := 0; i < 3; i++ {
		op <- i
	}
	close(op)
	wg.Wait()

	// Отдельный канал, при отправке в который, сигнализирует горутине об остановке
	wg.Add(1)
	q := make(chan int)
	go func() {
		fmt.Println("routine start")

		for {
			select {
			case <-q:
				fmt.Println("routine stops by receiving stop signal")
				wg.Done()
				return
			default:
				fmt.Println("routine working")
			}
		}
	}()

	q <- 1
	wg.Wait()

	//Отдельный канал, при закрытии которого, останавливается горутина
	wg.Add(1)
	stopChan := make(chan int)
	go func() {
		fmt.Println("routine started")
		for {
			select {
			case <-stopChan:
				fmt.Println("routine stops by closing a chan")
				wg.Done()
				return

			default:
				fmt.Println("routine working")
			}
		}
	}()

	close(stopChan)
	wg.Wait()

	//С помощью context. Так как пакет имеет функции для передачи сигналов отмены и тайм-аутов между горутинами
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func(c context.Context) {
		fmt.Println("routine start")
		for {
			select {
			case <-c.Done():
				fmt.Println("routine stops")
				wg.Done()
				return
			default:
				fmt.Println("routine working")
			}
		}
	}(ctx)
	cancel()
	wg.Wait()

	//завершение из-за условия
	wg.Add(1)
	a := 0
	go func() {
		fmt.Println("routine start")
		for a != 5 {
			a++
			fmt.Println("a= ", a)
		}
		fmt.Println("routine stop")
		wg.Done()
	}()
}
