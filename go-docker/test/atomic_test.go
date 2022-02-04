package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	wg := sync.WaitGroup{}

	var n uint64
	ch := p1(&n, &wg)
	p2(ch, &n, &wg)
	wg.Wait()
}

func p1(n *uint64, wg *sync.WaitGroup) chan int {
	ch := make(chan int)
	worker := func(workerId int, wg *sync.WaitGroup) {
		for i := 10; i <= 15; i++ {
			ch <- i
			fmt.Printf("### Worker:%d, Data sent:%d\n", workerId, i)
		}
		wg.Done()
		atomic.AddUint64(n, 1)
		fmt.Println("Done:", workerId)
	}

	const NumOfWorkers = 2
	for i := 1; i <= NumOfWorkers; i++ {
		wg.Add(1)
		go worker(i, wg)
	}
	return ch
}

func p2(ch chan int, n *uint64, wg *sync.WaitGroup) {
	worker := func(wg *sync.WaitGroup) {
		for {
			c := atomic.LoadUint64(n)
			fmt.Println("Count:", c)
			if c == 2 {
				return
			}
			select {
			case data := <-ch:
				fmt.Println(data)
			}
		}
	}

	go worker(wg)
}
