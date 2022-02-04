package test

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
)

const NumOfWorkers = 2

func TestAtomic(t *testing.T) {
	wg := sync.WaitGroup{}

	var doneWorkers uint64
	ch := p1(&doneWorkers, &wg)
	p2(ch, &doneWorkers, &wg)
	wg.Wait()
}

func p1(doneWorkers_p1 *uint64, wg *sync.WaitGroup) chan int {
	ch := make(chan int)
	worker := func(workerId int, doneWorkers *uint64, wg *sync.WaitGroup) {
		for d := 0; d <= rand.Intn(10); d++ {
			fmt.Printf("Worker:%d, Sending data:%d\n", workerId, d)
			ch <- d
		}
		wg.Done()
		atomic.AddUint64(doneWorkers, 1)
	}

	// Workers
	for i := 1; i <= NumOfWorkers; i++ {
		wg.Add(1)
		go worker(i, doneWorkers_p1, wg)
	}
	return ch
}

func p2(ch chan int, doneWorkers_p1 *uint64, wg *sync.WaitGroup) {
	worker := func(n *uint64, wg *sync.WaitGroup) {
		flag := true
		for flag {
			select {
			case d := <-ch:
				fmt.Printf("Received:%d\n", d)
			default:
				c := atomic.LoadUint64(n)
				if c == NumOfWorkers {
					fmt.Println("### All done:", c)
					flag = false
				}
			}

		}
	}

	go worker(doneWorkers_p1, wg)
}
