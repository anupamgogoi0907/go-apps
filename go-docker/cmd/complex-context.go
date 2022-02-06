package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

//func main() {
//	Init()
//}

type ComplexStage struct {
	ctx         context.Context
	wg          *sync.WaitGroup
	cancelFunc  context.CancelFunc
	noOfWorkers int
}

func Init() {
	wg := sync.WaitGroup{}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	const numOfWorkers = 2
	var doneCounters uint64

	data := a1(numOfWorkers, &doneCounters, cancel, ctx, &wg)
	a2(numOfWorkers, &doneCounters, data, cancel, ctx, &wg)
	wg.Wait()

}
func a1(numOfWorkers int, doneCounters *uint64, cancel context.CancelFunc, ctx context.Context, wg *sync.WaitGroup) chan int {
	data := make(chan int)
	worker := func(workerId int, cancel context.CancelFunc, ctx context.Context, wg *sync.WaitGroup) {
		for v := 0; v < 10; v++ {
			if v/2 == 1 {
				msg := "Error in Worker:" + strconv.Itoa(workerId) + " for value: " + strconv.Itoa(v)
				fmt.Println(msg)
				cancel()
				wg.Done()
				return
			} else {
				fmt.Printf(">>> Stage:%s, Worker:%d, Data:%d\n", "a1", workerId, v)
				data <- v
			}
		}
		wg.Done()
		atomic.AddUint64(doneCounters, 1)
	}

	for w := 1; w <= numOfWorkers; w++ {
		wg.Add(1)
		go worker(w, cancel, ctx, wg)
	}
	return data
}

func a2(numOfWorkers int, doneCounters *uint64, data chan int, cancel context.CancelFunc, ctx context.Context, wg *sync.WaitGroup) {
	worker := func(numOfWorkers int, doneCounters *uint64, data chan int, cancel context.CancelFunc, ctx context.Context, wg *sync.WaitGroup) {
		flag := true
		for flag {
			c := atomic.LoadUint64(doneCounters)
			select {
			case <-ctx.Done():
				fmt.Println("<<< Cancelled. Count", c)
				flag = false
				wg.Done()
			case d := <-data:
				fmt.Println("<<< Received:", d)
			default:
				if int(c) == numOfWorkers {
					fmt.Printf("<<< Count%d\n", c)
					flag = false
				}
			}
		}
		wg.Done()
	}

	wg.Add(1)
	go worker(numOfWorkers, doneCounters, data, cancel, ctx, wg)
}
