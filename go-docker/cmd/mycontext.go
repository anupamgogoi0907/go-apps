package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

var result = make(chan string)

func main() {
	wg := sync.WaitGroup{}

	out, count := processor1(&wg)
	processor2(out, &count, &wg)
	wg.Wait()
}

func processor1(wg *sync.WaitGroup) (chan int, uint64) {
	const NumSenders = 2
	ch := make(chan int)

	var countFinishedWorkers uint64
	// Worker
	worker := func(workerId int, ch chan int, wg *sync.WaitGroup) {
		n := rand.Intn(5)
		fmt.Printf("### Worker:%d, Total Values:%d\n", workerId, n)
		for i := 0; i < n; i++ {
			fmt.Printf("Sending value:%d\n", i)
			ch <- i
		}
		// Signal something that it has done processing.
		atomic.AddUint64(&countFinishedWorkers, 1)
		wg.Done()
	}

	// No of workers
	for i := 1; i <= NumSenders; i++ {
		wg.Add(1)
		go worker(i, ch, wg)
	}
	return ch, countFinishedWorkers
}

func processor2(ch chan int, count *uint64, wg *sync.WaitGroup) {
	consumer := func() {
		for {
			c := atomic.LoadUint64(count)
			fmt.Println("### Receiving #####", c)
			if c == 2 {
				return
			}
			select {
			case r := <-ch:
				fmt.Println("Received:", r)
			}
		}
	}
	go consumer()
}
