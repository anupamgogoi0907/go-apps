package pool

import (
	"fmt"
	"sync"
)

var chLines = make(chan []string, 5)
var chResults = make(chan []string)

type FileWorkerPool struct {
	FilePath           string
	NoOfLinesToProcess int
	NoOfWorkers        int
	waitGroup          *sync.WaitGroup
}

func (fw *FileWorkerPool) Run() {
	fw.waitGroup = &sync.WaitGroup{}
	go fw.consumeResults()
	fw.createWorkerPool()
	fw.ReadFileContent()
	fw.waitGroup.Wait()
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
	fw.waitGroup.Add(1)
	for i := 1; i <= 10; i++ {
		fmt.Println(<-chResults)
		fmt.Println("Results consumed.")
	}
	fw.waitGroup.Done()
}
