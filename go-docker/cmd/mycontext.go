package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

type Stage struct {
	noOfWorkers int
	doneWorkers *uint64
	ctx         context.Context
	cancelFunc  context.CancelFunc
	wg          *sync.WaitGroup
	data        chan int
	error       chan string
}

func main() {
	wg := sync.WaitGroup{}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	var n uint64
	s := &Stage{
		noOfWorkers: 2,
		doneWorkers: &n,
		ctx:         ctx,
		cancelFunc:  cancel,
		wg:          &wg,
	}
	IngestData(s)
	s1(s)
	wg.Wait()
}

func IngestData(s *Stage) {
	// Initialize the channel to send data.
	s.data = make(chan int)

	worker := func(workerId int, s *Stage) {
		n := rand.Intn(10)
		for d := 0; d <= n; d++ {
			fmt.Printf(">>> Stage:%s, Worker:%d, Data:%d/%d\n", "IngestData", workerId, d, n)
			s.data <- d
		}
		s.wg.Done()
		atomic.AddUint64(s.doneWorkers, 1)
	}

	for w := 1; w <= s.noOfWorkers; w++ {
		s.wg.Add(1)
		go worker(w, s)
	}

}

func s1(s *Stage) {
	worker := func(workerId int, s *Stage) {
		flag := true
		for flag {
			select {
			case d := <-s.data:
				fmt.Printf("<<< Stage:%s, Worker:%d, Data:%d\n", "s1", workerId, d)
			default:
				c := atomic.LoadUint64(s.doneWorkers)
				if int(c) == s.noOfWorkers {
					flag = false
					fmt.Println("<<< Received all in stage P1. Stopping P1.")
					s.wg.Done()
				}
			}
		}
	}
	s.wg.Add(1)
	go worker(1, s)
}

func s2(s *Stage) {

}
