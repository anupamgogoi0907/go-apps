package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	s1 := stage1([]int{1, 2, 3, 4, 5}, &wg)
	s2 := stage2(s1, &wg)
	s3 := stage3(s2, &wg)

	for d := range s3 {
		fmt.Println("Data receiver channel:", s3, "Data:", d)
	}
	wg.Wait()
}

func stage1(nums []int, wg *sync.WaitGroup) chan string {
	wg.Add(1)
	result := make(chan string)
	fmt.Println("Executing stage1.Data sender channel:", result)

	go func() {
		for _, n := range nums {
			fmt.Println("Receiving data in stage1.")
			v := strconv.Itoa(n) + "->stage1"
			result <- v
			fmt.Println("Sent data from stage1:", v)
		}
		close(result)
		wg.Done()
	}()
	return result
}

func stage2(in chan string, wg *sync.WaitGroup) chan string {
	result := make(chan string)
	fmt.Println("Executing stage2.Data receiver channel:", in, "Data sender channel:", result)
	wg.Add(1)

	go func(in chan string) {
		for n := range in {
			fmt.Println("Receiving data in stage2.Data receiver channel:", in)
			v := n + "->stage2"
			result <- v
			fmt.Println("Sent data from stage2:", v, "Data sender channel:", result)
		}
		close(result)
		wg.Done()
	}(in)

	return result
}

func stage3(in chan string, wg *sync.WaitGroup) chan string {
	result := make(chan string)
	fmt.Println("Executing stage3.Data receiver channel:", in, "Data sender channel:", result)
	wg.Add(1)

	go func(in chan string) {
		for n := range in {
			fmt.Println("Receiving data in stage3.Data receiver channel:", in)
			v := n + "->stage3"
			result <- v
			fmt.Println("Sent data from stage3:", v, "Data sender channel:", result)
		}
		close(result)
		wg.Done()
	}(in)

	return result
}
