package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	InitSimpleContext()
}
func InitSimpleContext() {
	wg := sync.WaitGroup{}
	ctxRoot := context.Background()
	ctxCancel, cancel := context.WithCancel(ctxRoot)

	data := make(chan int)
	b1(data, cancel, ctxCancel, &wg)
	b2(data, cancel, ctxCancel, &wg)
	wg.Wait()
}

func b1(data chan int, cancel context.CancelFunc, ctxCancel context.Context, wg *sync.WaitGroup) {
	worker := func(workerId int, data chan int, ctxCancel context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		flag := true
		d := 0
		for flag {
			if d == 5 {
				flag = false
				cancel()
			} else {
				d = d + 1
				fmt.Printf(">>> B1, Worker:%d, Received:%d\n", workerId, d)
				data <- d
			}
		}
	}

	wg.Add(1)
	go worker(1, data, ctxCancel, wg)
}
func b2(data chan int, cancel context.CancelFunc, ctxCancel context.Context, wg *sync.WaitGroup) {
	worker := func(workerId int, data chan int, ctxCancel context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		flag := true
		for flag {
			select {
			case <-ctxCancel.Done():
				fmt.Println("B2 exited.")
				flag = false
			case d := <-data:
				fmt.Printf("<<< B2, Worker:%d, Received:%d\n", workerId, d)
			default:

			}
		}
	}

	wg.Add(2)
	go worker(1, data, ctxCancel, wg)
	go worker(2, data, ctxCancel, wg)
}
