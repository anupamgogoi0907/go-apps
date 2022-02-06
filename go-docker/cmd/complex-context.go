package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	InitComplexContext()
}
func InitComplexContext() {
	wg := sync.WaitGroup{}
	ctxRoot := context.Background()

	data := make(chan int)

	// Stage 1
	ctxCancelStage1, cancel := context.WithCancel(ctxRoot)
	a1(data, cancel, ctxCancelStage1, &wg)

	// Stage 2
	ctxCancelStage2, _ := context.WithCancel(ctxCancelStage1)
	a2(data, ctxCancelStage2, &wg)

	wg.Wait()
}
func a1(data chan int, cancel context.CancelFunc, ctxCancel context.Context, wg *sync.WaitGroup) {
	worker := func(workerId int, data chan int, cancel context.CancelFunc, ctxCancel context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		flag := true

		// Data block.
		d := 0
		for flag {
			if d == 5 {
				flag = false
				cancel()
			} else {
				d = d + 1
				fmt.Printf(">>> A1, Worker:%d, Sent:%d\n", workerId, d)
				data <- d
			}
		}
	}

	// Invoke workers.
	wg.Add(1)
	go worker(1, data, cancel, ctxCancel, wg)
}
func a2(data chan int, ctxCancel context.Context, wg *sync.WaitGroup) {
	worker := func(workerId int, data chan int, ctxCancel context.Context, wg *sync.WaitGroup) {
		defer wg.Done()

		// Data block.
		flag := true
		for flag {
			select {
			case <-ctxCancel.Done():
				fmt.Println("### A2 exited.")
				flag = false
			case d := <-data:
				fmt.Printf("<<< A2, Worker:%d, Received:%d\n", workerId, d)
			default:

			}
		}
	}

	wg.Add(1)
	go worker(1, data, ctxCancel, wg)
}
