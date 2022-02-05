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

//func main() {
//	InitPipeline()
//}
func InitPipeline() {
	wg := sync.WaitGroup{}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// Source
	var n uint64
	stageSource := &Stage{
		noOfWorkers: 2,
		doneWorkers: &n,
		ctx:         ctx,
		cancelFunc:  cancel,
		wg:          &wg,
		data:        make(chan int),
		error:       make(chan string),
	}
	IngestData(stageSource)

	// Stagex 1
	var n1 uint64
	stage1 := &Stage{
		noOfWorkers: 2,
		doneWorkers: &n1,
		ctx:         ctx,
		cancelFunc:  cancel,
		wg:          &wg,
		data:        make(chan int),
		error:       make(chan string),
	}
	s1(stageSource, stage1)

	// Stagex 2
	var n2 uint64
	stage2 := &Stage{
		noOfWorkers: 2,
		doneWorkers: &n2,
		ctx:         ctx,
		cancelFunc:  cancel,
		wg:          &wg,
		data:        make(chan int),
		error:       make(chan string),
	}
	s2(stage1, stage2)
	wg.Wait()
}

func IngestData(cur *Stage) {
	worker := func(workerId int, s *Stage) {
		n := rand.Intn(10)
		for d := 0; d <= n; d++ {
			fmt.Printf(">>> Stagex:%s, Worker:%d, Data:%d/%d\n", "IngestData", workerId, d, n)
			s.data <- d
		}
		s.wg.Done()
		atomic.AddUint64(s.doneWorkers, 1)
	}

	for w := 1; w <= cur.noOfWorkers; w++ {
		cur.wg.Add(1)
		go worker(w, cur)
	}

}

func s1(prev *Stage, cur *Stage) {
	worker := func(workerId int, prev *Stage, cur *Stage) {
		flag := true
		for flag {
			select {
			case d := <-prev.data:
				fmt.Printf("<<< Stagex:%s, Worker:%d, Data:%d\n", "s1", workerId, d)
				cur.data <- d * 10
			default:
				c := atomic.LoadUint64(prev.doneWorkers)
				if int(c) == prev.noOfWorkers {
					flag = false
					fmt.Println("### Received all in stage s1. Stopping s1.")
					atomic.AddUint64(cur.doneWorkers, 1)
					cur.wg.Done()
				}
			}
		}
	}
	for w := 1; w <= cur.noOfWorkers; w++ {
		cur.wg.Add(1)
		go worker(w, prev, cur)
	}
}

func s2(prev *Stage, cur *Stage) {
	worker := func(workerId int, prev *Stage, cur *Stage) {
		flag := true
		for flag {
			select {
			case d := <-prev.data:
				fmt.Printf("<<< Stagex:%s, Worker:%d, Data:%d\n", "s2", workerId, d)
			default:
				c := atomic.LoadUint64(prev.doneWorkers)
				if int(c) == prev.noOfWorkers {
					flag = false
					fmt.Println("### Received all in stage s2. Stopping s2.")
					cur.wg.Done()
				}
			}
		}
	}
	for w := 1; w <= cur.noOfWorkers; w++ {
		cur.wg.Add(1)
		go worker(w, prev, cur)
	}
}
