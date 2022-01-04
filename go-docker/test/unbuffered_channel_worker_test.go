package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var jobs1 = make(chan int, 5)
var result1 = make(chan int)

func TestPoolWithUnbufferedResultChannel(t *testing.T) {
	wg := sync.WaitGroup{}

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second * 2)
		jobs1 <- 10 * i
		fmt.Println(<-result1)
	}
	close(jobs1)

	wg.Wait()
}

func worker(id int, wg *sync.WaitGroup) {
	for j := range jobs1 {
		fmt.Println("Worker: ", id, " processed:", j)
		result1 <- 10 * j
		fmt.Println("Results sent")
	}
	wg.Done()
}

func Test(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Consumer
	//go consumer(&wg)
	//jobs1 <- 10

	// Producer
	go producer(&wg)
	fmt.Println(<-jobs1)
	wg.Wait()
}

func consumer(wg *sync.WaitGroup) {
	fmt.Println(<-jobs1)
	wg.Done()
}

func producer(wg *sync.WaitGroup) {
	jobs1 <- 100
	wg.Done()
}
