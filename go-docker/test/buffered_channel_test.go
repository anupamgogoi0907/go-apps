package test

import (
	"fmt"
	"testing"
	"time"
)

// Worker
func Worker(id int, jobs chan int, result chan int) {
	for j := range jobs {
		time.Sleep(time.Second * 2)
		fmt.Println("Worker: ", id, ", Jobs done: ", j)
		result <- j * 1000
	}
}

func PrintResult(results chan int) {
	for {
		fmt.Println(<-results)
	}
}

func TestBufferedChannel(t *testing.T) {
	const numJobs = 3
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	go PrintResult(results)

	// Below workers (threads) will read data from the same channel(queue) i.e. jobs and put the results to the same channel results.
	go Worker(1, jobs, results)
	go Worker(2, jobs, results)

	for i := 1; i <= numJobs; i++ {
		jobs <- i * 10
	}
	close(jobs)

	time.Sleep(time.Minute)
}
