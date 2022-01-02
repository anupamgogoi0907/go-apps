package test

import (
	"fmt"
	"sync"
	"testing"
)

var jobs = make(chan int, 10)
var results = make(chan int, 10)

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := 1000 * job
		results <- output
	}
	wg.Done()
}
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		jobs <- i
	}
	close(jobs)
}
func result() {
	for result := range results {
		fmt.Println(result)
	}
}
func TestWP(t *testing.T) {

	noOfJobs := 1000000
	go allocate(noOfJobs)

	go result()
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)

}
