package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test(t *testing.T) {
	wg := sync.WaitGroup{}
	noOfJobs := 50
	noOfWorkers := 2
	// We need a goroutine to consume (read) results for channel results.
	go consumeResults(noOfJobs, &wg)
	wg.Add(1)

	// We are creating a fixed no (i.e 2 from noOfWorkers variable) of workers (goroutines).
	for i := 1; i <= noOfWorkers; i++ {
		wg.Add(1)
		go produceResults(i, &wg)
	}

	// We are sending 50 tasks to the channel jobs.
	for i := 1; i <= noOfJobs; i++ {
		time.Sleep(time.Second)
		// Send a task.(In this case just an integer).
		// Below call is not blocking until we reach the size of the channel jobs i.e. 5
		// Note that we are sending 50 tasks to the jobs channel but the buffer size of it is only 5. Why it did not enter into deadlock state?
		// The answer is, we have the 2 produceResults goroutines reading and emptying the jobs channel concurrently.
		jobs <- i
	}
	// Close the channel jobs after tasks are sent (written) to it.
	close(jobs)

	wg.Wait()
}

// The produceResults goroutine. There will be two concurrent goroutines reading (consuming/emptying) values from the jobs channel.
func produceResults(id int, wg *sync.WaitGroup) {
	for j := range jobs {
		fmt.Println("Worker:", id, "is processing:", j)
		// After reading the value from the jobs channel it will put the results in the results channel.
		// The below code will be blocking until someone (goroutine) reads value from it.
		// To do that, we have a concurrent goroutine called consumeResults is already running.
		results <- 10 * j
		fmt.Println("Results produced.")
	}
	wg.Done()
}

// It will consume the values sent to the results channel. The results are produced by the produceResults function concurrently.
// Note that we have 2 workers running concurrently de
func consumeResults(noOfJobs int, wg *sync.WaitGroup) {
	for i := 1; i <= noOfJobs; i++ {
		// Below line blocks the for loop until someone sends (writes) values to the results channel.
		fmt.Println(<-results)
		fmt.Println("Results consumed.")
	}
	wg.Done()
}
