package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var result = make(chan string)

func main() {
	wg := sync.WaitGroup{}

	out, done := processor1(&wg)
	processor2(out, done, &wg)
	wg.Wait()
}

func processor1(wg *sync.WaitGroup) (chan int, chan bool) {
	const NumSenders = 2
	ch := make(chan int)
	done := make(chan bool, NumSenders)
	// Worker
	worker := func(workerId int, out chan int, wg *sync.WaitGroup) {
		n := rand.Intn(5)
		fmt.Printf("### Worker:%d, Total Values:%d\n", workerId, n)
		for i := 0; i < n; i++ {
			fmt.Printf("Value:%d\n", i)
			out <- i
		}
		// Signal something that it has done processing.
		//done <- true
		wg.Done()
	}

	// No of workers
	for i := 1; i <= NumSenders; i++ {
		wg.Add(1)
		go worker(i, ch, wg)
	}
	return ch, done
}

func processor2(ch chan int, done chan bool, wg *sync.WaitGroup) {
	consumer := func() {
		for {
			select {
			case d := <-ch:
				fmt.Println("RECEIVED:", d)
			}
			//wg.Done()
		}
	}
	//wg.Add(1)
	go consumer()
}
