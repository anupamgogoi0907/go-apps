package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	s1 := stage1([]int{1, 2, 3}, &wg)
	s2 := stage2(s1, &wg)
	s3 := stage3(s2, &wg)

	for d := range s3 {
		fmt.Println("### Final Source channel:", s3, ",Data:", d)
	}
	wg.Wait()
}

func stage1(nums []int, wg *sync.WaitGroup) chan string {
	wg.Add(1)
	result := make(chan string)

	worker := func(workerId int) {
		for _, n := range nums {
			fmt.Printf("RECEIVING:Stage1,Worker:%d,Data:%d,Source Channel:%x,Target Channel:%x\n", workerId, n, nil, result)
			v := strconv.Itoa(n) + "->stage1"
			result <- v
			fmt.Printf("SENT:Stage1,Worker:%d,Data:%s,Target Channel:%x\n", workerId, v, result)
		}
		close(result)
		wg.Done()
	}
	go worker(1)
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
