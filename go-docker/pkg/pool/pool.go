package pool

import (
	"fmt"
	"github.com/anupamgogoi0907/go-apps/go-docker/pkg/model"
	"io/ioutil"
	"log"
	"sync"
)

var jobs chan model.Job
var results chan model.Result

type WorkerPool struct {
	CurDir      string
	NoOfWorkers int
	waitGroup   *sync.WaitGroup
}

// configurePool Configure the channels.
func (wp *WorkerPool) configurePool(bufferSize int) {
	jobs = make(chan model.Job, bufferSize)
	results = make(chan model.Result, bufferSize)
}

// Run the main point of entry to the code.
func (wp *WorkerPool) Run() {
	wp.waitGroup = &sync.WaitGroup{}
	files := GetFilesOfCurrentDirectory(wp.CurDir)
	wp.configurePool(len(files))
	wp.allocateJobs(files)
	wp.createWorkerPool()
	wp.result()
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
	wp.waitGroup.Wait()
	close(results)
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
func (wp *WorkerPool) result() {
	for result := range results {
		fmt.Println("Processed file: ", result.LogLines)
	}
}

// GetFilesOfCurrentDirectory finds the files in the provided directory dir
func GetFilesOfCurrentDirectory(dir string) (arr []string) {
	if dir == "" {
		log.Panicln("Empty string")
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Panicln(err)
	}

	var arrFiles []string
	for _, file := range files {
		arrFiles = append(arrFiles, dir+file.Name())
	}
	return arrFiles
}
