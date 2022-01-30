package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var result = make(chan string)

func main() {
	wg := sync.WaitGroup{}

	out := processor1(&wg)
	processor2(out, &wg)
	wg.Wait()
}

func processor1(wg *sync.WaitGroup) chan int {
	ch := make(chan int)

	// Worker
	worker := func(id int, out chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		v := rand.Intn(5)
		for i := 0; i < v; i++ {
			fmt.Printf("SENDING:Channel::%x, Worker:%d, Value:%d\n", out, i, v)
			out <- i
		}
	}

	// No of workers
	const NumSenders = 2
	for i := 1; i <= NumSenders; i++ {
		wg.Add(1)
		go worker(i, ch, wg)
	}
	return ch
}

func processor2(ch chan int, wg *sync.WaitGroup) {
	select {
	case data := <-ch:
		fmt.Println("RECEIVING::", data)
	}
}
