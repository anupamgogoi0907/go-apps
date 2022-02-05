package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

type Stage struct {
	noOfWorkers     int
	noOfDoneWorkers *uint64
	ctx             context.Context
	cancelFunc      context.CancelFunc
	wg              *sync.WaitGroup
}

func main() {
	wg := sync.WaitGroup{}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	var n uint64
	s := &Stage{
		noOfWorkers:     2,
		noOfDoneWorkers: &n,
		ctx:             ctx,
		cancelFunc:      cancel,
		wg:              &wg,
	}
	IngestData(s)
	p1(nil, s)
	wg.Wait()
}

func IngestData(s *Stage) {
	worker := func(workerId int, s *Stage) {
		fmt.Printf(">>> Worker:%d\n", workerId)
		s.wg.Done()
		atomic.AddUint64(s.noOfDoneWorkers, 1)
	}

	for w := 1; w <= s.noOfWorkers; w++ {
		s.wg.Add(1)
		go worker(w, s)
	}

}

func p1(ch chan int, s *Stage) {
	worker := func(workerId int, s *Stage) {
		flag := true
		for flag {
			c := atomic.LoadUint64(s.noOfDoneWorkers)
			if int(c) == s.noOfWorkers {
				flag = false
				fmt.Println("<<< Received in P1")
				s.wg.Done()
				return
			}
		}
	}
	s.wg.Add(1)
	go worker(1, s)
}
