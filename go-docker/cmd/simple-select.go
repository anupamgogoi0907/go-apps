package main

import (
	"fmt"
	"sync"
)

var (
	data = make(chan int)
)

//func main() {
//	InitSimpleSelect()
//}

func InitSimpleSelect() {
	wg := sync.WaitGroup{}

	c1(&wg)
	c2(&wg)

	wg.Wait()
}

func c1(wg *sync.WaitGroup) {
	worker := func(workerId int, wg *sync.WaitGroup) {
		defer wg.Done()
		data <- 10
		data <- 20
	}

	wg.Add(1)
	go worker(1, wg)
}
func c2(wg *sync.WaitGroup) {
	worker := func(workerId int, wg *sync.WaitGroup) {
		defer wg.Done()
		flag := true
		for flag {
			select {
			case d := <-data:
				fmt.Printf("<<< C2, Worker:%d, Received:%d\n", workerId, d)
			default:
				fmt.Printf("<<< C2, Worker:%d\n", workerId)
			}
		}
	}

	wg.Add(1)
	go worker(1, wg)
}
