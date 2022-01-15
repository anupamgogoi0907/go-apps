package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	pipeline([]int{1, 2}, &wg)
	wg.Wait()
}

func pipeline(data []int, wg *sync.WaitGroup) {
	wg.Add(1)
	in := make(chan string)
	go func() {
		for _, n := range data {
			v := strconv.Itoa(n)
			in <- v
		}
		close(in)
		wg.Done()
	}()

	result1 := stage1(in, wg)
	result2 := stage2(result1, wg)
	result3 := stage3(result2, wg)
	finish(result3)
}

func stage1(in chan string, wg *sync.WaitGroup) chan string {
	result := make(chan string)
	wg.Add(1)

	worker := func(workerId int, in chan string) {
		for n := range in {
			fmt.Printf("RECEIVING.Stage1,Worker:%d,Data:%s,Source Channel:%x,Target Channel:%x\n", workerId, n, in, result)
			v := n + "->stage1"
			result <- v
			fmt.Printf("SENT:Stage1,Worker:%d,Data:%s,Target Channel:%x\n", workerId, v, result)
		}
		close(result)
		wg.Done()
	}
	go worker(1, in)

	return result
}

func stage2(in chan string, wg *sync.WaitGroup) chan string {
	result := make(chan string)
	wg.Add(1)

	worker := func(workerId int, in chan string) {
		for n := range in {
			fmt.Printf("RECEIVING.Stage2,Worker:%d,Data:%s,Source Channel:%x,Target Channel:%x\n", workerId, n, in, result)
			v := n + "->stage2"
			result <- v
			fmt.Printf("SENT:Stage2,Worker:%d,Data:%s,Target Channel:%x\n", workerId, v, result)
		}
		close(result)
		wg.Done()
	}
	go worker(1, in)

	return result
}

func stage3(in chan string, wg *sync.WaitGroup) chan string {
	result := make(chan string)
	wg.Add(1)

	worker := func(workerId int, in chan string) {
		for n := range in {
			fmt.Printf("RECEIVING.Stage3,Worker:%d,Data:%s,Source Channel:%x,Target Channel:%x\n", workerId, n, in, result)
			v := n + "->stage3"
			result <- v
			fmt.Printf("SENT:Stage3,Worker:%d,Data:%s,Target Channel:%x\n", workerId, v, result)
		}
		close(result)
		wg.Done()
	}
	go worker(1, in)
	return result
}

func finish(in chan string) {
	for d := range in {
		fmt.Println("### Final Source channel:", in, ",Data:", d)
	}
}
