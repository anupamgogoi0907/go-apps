package test

import (
	"fmt"
	"sync"
	"testing"
)

var jobs = make(chan int, 5)
var results = make(chan int, 5)

type MyWorker struct {
	noOfJobs    int
	noOfWorkers int
	waitGroup   *sync.WaitGroup
}

func (s MyWorker) worker(workerId int) {
	for job := range jobs {
		fmt.Println("Worker:", workerId, ", Job: ", job)
		output := 1000 * job
		results <- output
	}
	s.waitGroup.Done()
}
func (s MyWorker) createWorkerPool() {
	for i := 1; i <= s.noOfWorkers; i++ {
		s.waitGroup.Add(1)
		go s.worker(i)
	}
	s.waitGroup.Wait()
	close(results)
}
func (s MyWorker) allocate() {
	for i := 1; i <= s.noOfJobs; i++ {
		jobs <- i
	}
	close(jobs)
}

func (s MyWorker) result() {
	for result := range results {
		fmt.Println(result)
	}
}

func TestWP(t *testing.T) {
	s := MyWorker{
		noOfJobs:    10,
		noOfWorkers: 2,
		waitGroup:   &sync.WaitGroup{},
	}
	go s.allocate()
	go s.result()
	//s.createWorkerPool()

}
