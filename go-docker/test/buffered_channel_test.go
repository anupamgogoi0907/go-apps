package test

import (
	"fmt"
	"testing"
	"time"
)

// TestMultipleWorker
func TestMultipleWorker(t *testing.T) {
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

// TestSingleWorker
func TestSingleWorker(t *testing.T) {
	jobs := make(chan int, 2)

	go TempWorker(1, jobs)

	jobs <- 10
	jobs <- 20
	jobs <- 20

	time.Sleep(time.Minute)
}

func TempWorker(id int, jobs chan int) {
	for {
		fmt.Println("Worker: ", id, " ,Jobs: ", <-jobs)
	}
}
