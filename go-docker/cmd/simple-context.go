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
	worker := func(data chan int, ctxCancel context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		flag := true
		d := 0
		for flag {
			if d == 5 {
				flag = false
				cancel()
			} else {
				d = d + 1
				data <- d
			}
		}
	}

	wg.Add(1)
	go worker(data, ctxCancel, wg)
}
func b2(data chan int, cancel context.CancelFunc, ctxCancel context.Context, wg *sync.WaitGroup) {
	worker := func(data chan int, ctxCancel context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		flag := true
		for flag {
			select {
			case <-ctxCancel.Done():
				fmt.Println("B2 exited.")
				flag = false
			case d := <-data:
				fmt.Printf("<<< B2, Received:%d\n", d)
			default:

			}
		}
	}

	wg.Add(1)
	go worker(data, ctxCancel, wg)
}
