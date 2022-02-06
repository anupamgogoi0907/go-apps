package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

func main() {
	InitErrorHandling()
}

var (
	error = make(chan string)

	b1_noWorkers   int = 2
	b1_doneWorkers uint64
	b1_data        = make(chan int)

	b2_noWorkers   int = 1
	b2_doneWorkers uint64
)

func InitErrorHandling() {
	wg := sync.WaitGroup{}
	b1(&wg)
	b2(&wg)
	wg.Wait()
}

func b1(wg *sync.WaitGroup) {
	worker := func(workerId int, wg *sync.WaitGroup) {
		defer wg.Done()
		for d := 1; d <= 5; d++ {
			n := rand.Intn(10)
			if d == n {
				e := "Some error"
				fmt.Printf(">>> B1, Worker:%d, Sent:%s\n", workerId, e)
				error <- e
			}
			fmt.Printf(">>> B1, Worker:%d, Sent:%d\n", workerId, d)
			b1_data <- d
		}
		atomic.AddUint64(&b1_doneWorkers, 1)
	}

	// Workers
	for w := 1; w <= b1_noWorkers; w++ {
		wg.Add(1)
		go worker(w, wg)
	}
}

func b2(wg *sync.WaitGroup) {
	worker := func(workerId int, wg *sync.WaitGroup) {
		defer wg.Done()
		flag := true
		for flag {
			select {
			case d := <-b1_data:
				fmt.Printf(">>> B2, Worker:%d, Received:%d\n", workerId, d)
			case e := <-error:
				fmt.Printf(">>> B2, Worker:%d, Received:%s\n", workerId, e)
			default:
				c := atomic.LoadUint64(&b1_doneWorkers)
				if int(c) == b1_noWorkers {
					flag = false
					fmt.Printf(">>> B2 has received all.")
				}
			}
		}
	}

	// Workers
	for w := 1; w <= b2_noWorkers; w++ {
		wg.Add(1)
		go worker(w, wg)
	}
}
