package pool

import (
	"fmt"
	"github.com/anupamgogoi0907/go-apps/go-docker/pkg/model"
	"github.com/anupamgogoi0907/go-apps/go-docker/pkg/utility"
	"sync"
)

var jobs chan model.Job
var results chan model.Result

type WorkerPool struct {
	CurDir      string
	NoOfWorkers int
	waitGroup   *sync.WaitGroup
}

// Run the main point of entry to the code.
func (wp *WorkerPool) Run() {
	wp.waitGroup = &sync.WaitGroup{}
	files := utility.GetFilesOfCurrentDirectory(wp.CurDir)

	go wp.result(len(files))
	wp.configurePool(len(files))
	wp.allocateJobs(files)
	wp.createWorkerPool()

	wp.waitGroup.Wait()
}

// configurePool Configure the channels.
func (wp *WorkerPool) configurePool(bufferSize int) {
	jobs = make(chan model.Job, bufferSize)
	results = make(chan model.Result)
}

// allocateJobs allocates the jobs i.e. each log file is queued in the jobs channel
func (wp *WorkerPool) allocateJobs(files []string) {
	for _, file := range files {
		fmt.Println("Allocating file: ", file)
		job := model.Job{
			FilePath: file,
		}
		jobs <- job
	}
	close(jobs)
}

// createWorkerPool it creates the provided no of workers.
func (wp *WorkerPool) createWorkerPool() {
	for i := 0; i < wp.NoOfWorkers; i++ {
		wp.waitGroup.Add(1)
		go wp.worker()
	}
}

func (wp *WorkerPool) worker() {
	for job := range jobs {
		fmt.Println("Processing file: ", job.FilePath)
		res := model.Result{
			LogLines: job.FilePath,
		}
		results <- res
	}
	wp.waitGroup.Done()
}
func (wp *WorkerPool) result(noOfResults int) {
	wp.waitGroup.Add(1)
	for i := 0; i < noOfResults; i++ {
		fmt.Println(<-results)
	}
	wp.waitGroup.Done()
}
