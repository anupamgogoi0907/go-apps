package pool

import (
	"bufio"
	"log"
	"os"
	"sync"
)

var chLines = make(chan []string, 5)

type FileWorkerPool struct {
	FilePath           string
	NoOfLinesToProcess int
	NoOfWorkers        int
	TargetLocation     string
	waitGroup          *sync.WaitGroup
}

func (fw *FileWorkerPool) Run() {
	// 0. Validate input.
	fw.validate()

	// 1. Assign workers.
	fw.createWorkerPool()

	// 2. Read large file
	fw.readFileContent()

	fw.waitGroup.Wait()
}

func (fw *FileWorkerPool) validate() {
	if fw.waitGroup == nil {
		fw.waitGroup = &sync.WaitGroup{}
	}
	if fw.TargetLocation == "" {
		pwd, _ := os.Getwd()
		fw.TargetLocation = pwd
	}
	if fw.NoOfWorkers == 0 {
		fw.NoOfWorkers = 2
	}
	if fw.NoOfLinesToProcess == 0 {
		fw.NoOfLinesToProcess = 10000
	}
}

// Create a worker pool.
func (fw *FileWorkerPool) createWorkerPool() {
	for w := 1; w <= fw.NoOfWorkers; w++ {
		fw.waitGroup.Add(1)
		go fw.worker(w)
	}
}

// Main processing logic goes here.
func (fw *FileWorkerPool) worker(workerId int) {
	for lines := range chLines {
		log.Println("Worker", workerId, "is processing")
		log.Println("Worker", workerId, "has processed", lines)
	}
	fw.waitGroup.Done()
}

// Read file
func (fw *FileWorkerPool) readFileContent() {
	if fw.FilePath == "" {
		log.Panicln("No file provided.")
	}
	file, err := os.Open(fw.FilePath)
	if err != nil {
		log.Panicln(err)
	}
	scanner := bufio.NewScanner(file)

	// Count the number of batches of lines sent to chLines
	lineBuffer := []string{}
	for scanner.Scan() {
		if len(lineBuffer) >= fw.NoOfLinesToProcess {
			log.Println("Send lineBuffer to chLines to be processed.")
			chLines <- lineBuffer
			lineBuffer = []string{}
		}
		line := scanner.Text()
		lineBuffer = append(lineBuffer, line)

	}
	// Check if there is remaining lines in the lineBuffer.
	if len(lineBuffer) != 0 {
		lineBuffer = nil
	}
	close(chLines)
	defer file.Close()
}
