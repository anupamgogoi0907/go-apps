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
	worker := func(wg *sync.WaitGroup) {
		ch <- 10
		wg.Done()
		atomic.AddUint64(n, 10000)
		fmt.Println("Done")
	}

	wg.Add(1)
	go worker(wg)
	return ch
}

func p2(ch chan int, n *uint64, wg *sync.WaitGroup) {
	worker := func(wg *sync.WaitGroup) {
		<-ch
		fmt.Println(atomic.LoadUint64(n))
		wg.Done()
	}

	wg.Add(1)
	go worker(wg)
}
