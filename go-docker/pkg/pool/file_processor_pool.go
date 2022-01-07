package pool

import (
	"fmt"
	"sync"
)

var chLines = make(chan []string, 5)
var chResults = make(chan []string, 5)

type FileWorkerPool struct {
	FilePath           string
	NoOfLinesToProcess int
	NoOfWorkers        int
	waitGroup          *sync.WaitGroup
}

func (fw *FileWorkerPool) Run() {
	fw.waitGroup = &sync.WaitGroup{}
	// 1. Open the consumer
	go fw.consumeResults()

	// 2. Create the workers to read data from chLines and send results to chResults
	fw.createWorkerPool()

	// 3. Read file and send content to chLines
	fw.ReadFileContent()

	// 4. Wait for all goroutines to finish
	fw.waitGroup.Wait()

	// 5. Close the chResults
	close(chResults)
}
func (fw *FileWorkerPool) createWorkerPool() {
	for i := 1; i <= fw.NoOfWorkers; i++ {
		go fw.produceResults(i)
	}
}
func (fw *FileWorkerPool) ReadFileContent() {
	for i := 1; i <= 10; i++ {
		chLines <- []string{"Hello"}
	}
	close(chLines)
}

func (fw *FileWorkerPool) produceResults(id int) {
	fw.waitGroup.Add(1)
	for j := range chLines {
		fmt.Println("Worker:", id, "is processing:", j)
		chResults <- j
		fmt.Println("Worker:", id, "has processed:", j)
	}
	fw.waitGroup.Done()
}

func (fw *FileWorkerPool) consumeResults() {
	for r := range chResults {
		fmt.Println(r)
		fmt.Println("Results consumed.")
	}

}
