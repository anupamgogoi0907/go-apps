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
		for d := 1; d <= 3; d++ {
			ch <- d
		}
		wg.Done()
		atomic.AddUint64(n, 1)
	}

	// Workers
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, wg)
	}
	return ch
}

func p2(ch chan int, n *uint64, wg *sync.WaitGroup) {
	worker := func(n *uint64, wg *sync.WaitGroup) {
		flag := true
		for flag {
			select {
			case d := <-ch:
				fmt.Println(d)
			default:
				c := atomic.LoadUint64(n)
				if c == 2 {
					fmt.Println("### All done:", c)
					flag = false
				}
			}

		}
	}

	go worker(n, wg)
}
