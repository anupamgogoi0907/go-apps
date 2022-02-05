package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
)

var result = make(chan string)

func main() {
	wg := sync.WaitGroup{}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	var doneWorkers_p1 uint64
	ch, err := processor1(&doneWorkers_p1, ctx, &wg)
	processor2(ch, err, &doneWorkers_p1, ctx, &wg)

	wg.Wait()
	cancel()
}

func processor1(doneWorkers_p1 *uint64, ctx context.Context, wg *sync.WaitGroup) (chan int, chan string) {
	ch := make(chan int)
	err := make(chan string)

	worker := func(workerId int, wg *sync.WaitGroup) {
		n := rand.Intn(50)
		for v := 0; v < n; v++ {
			if v/2 == 1 {
				msg := "Error in Worker:" + strconv.Itoa(workerId) + " for value: " + strconv.Itoa(v)
				err <- msg
				wg.Done()
				atomic.AddUint64(doneWorkers_p1, 1)
				return
			} else {
				fmt.Printf(">>> Worker:%d produced:%d\n", workerId, v)
				ch <- v
			}

		}
		wg.Done()
		atomic.AddUint64(doneWorkers_p1, 1)
	}

	// Workers
	for w := 1; w <= 2; w++ {
		wg.Add(1)
		go worker(w, wg)
	}

	return ch, err
}

func processor2(ch chan int, err chan string, doneWorkers_p1 *uint64, ctx context.Context, wg *sync.WaitGroup) {
	worker := func(workerId int, wg *sync.WaitGroup) {
		flag := true
		for flag {
			select {
			case e := <-err:
				fmt.Printf("<<< Received by Worker:%d, Value:%s\n", workerId, e)
			case d := <-ch:
				fmt.Printf("<<< Received by Worker:%d, Value:%d\n", workerId, d)
			default:
				c := atomic.LoadUint64(doneWorkers_p1)
				if c == 2 {
					flag = false
				}
			}
		}
	}
	go worker(2, wg)
	go worker(1, wg)
}
