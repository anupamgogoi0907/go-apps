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
	fmt.Println("Executing stage1.Target channel:", result)

	worker := func(workerId int) {
		for _, n := range nums {
			fmt.Println("Receiving data in stage1,WorkerID", workerId, "Source channel:", nil)
			v := strconv.Itoa(n) + "->stage1"
			result <- v
			fmt.Println("Sent data from stage1:", v, "Target channel:", result)
		}
		close(result)
		wg.Done()
	}
	go worker(1)
	return result
}

func stage2(in chan string, wg *sync.WaitGroup) chan string {
	result := make(chan string)
	fmt.Println("Executing stage2.Source channel:", in, ",Target channel:", result)
	wg.Add(1)

	worker := func(workerId int, in chan string) {
		for n := range in {
			fmt.Println("Receiving data in stage2,WorkerID", workerId, "Source channel:", nil)
			v := n + "->stage2"
			result <- v
			fmt.Println("Sent data from stage2:", v, ",Target channel:", result)
		}
		close(result)
		wg.Done()
	}
	go worker(1, in)

	return result
}

func stage3(in chan string, wg *sync.WaitGroup) chan string {
	result := make(chan string)
	fmt.Println("Executing stage3.Source channel:", in, ",Target channel:", result)
	wg.Add(1)

	worker := func(workerId int, in chan string) {
		for n := range in {
			fmt.Println("Receiving data in stage3,WorkerID", workerId, "Source channel:", nil)
			v := n + "->stage3"
			result <- v
			fmt.Println("Sent data from stage3:", v, ",Target channel:", result)
		}
		close(result)
		wg.Done()
	}
	go worker(1, in)
	return result
}
