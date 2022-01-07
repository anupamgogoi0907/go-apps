package pool

import (
	"bufio"
	"log"
	"os"
	"sync"
)

var chLines = make(chan []string, 5)
var chResults = make(chan []string, 5)

type FileWorkerPool struct {
	FilePath           string
	NoOfLinesToProcess int
	waitGroup          *sync.WaitGroup
}

func (fw *FileWorkerPool) Run() {
	fw.waitGroup = &sync.WaitGroup{}

	// Allocate a goroutine to consume results.
	go fw.consumeResultsWorker()

	// Allocate 10 workers to process file lines.
	noOfWorkers := 2
	for i := 1; i <= noOfWorkers; i++ {
		fw.waitGroup.Add(1)
		go fw.produceResultsWorker(i)
	}
	fw.ReadFileContent()

	fw.waitGroup.Wait()
}

// 10 produceResultsWorker are working concurrently to process data sent to chLines channel.
func (fw *FileWorkerPool) produceResultsWorker(workerId int) {
	for l := range chLines {
		// Add processing code.
		log.Println("Worker:", workerId, "is processing.", len(l))
		//time.Sleep(time.Second * 2)
		// Below call will block the for loop until someone processes the result.
		chResults <- []string{"Processed"}
		log.Println("Lines has been processed.")
	}
	fw.waitGroup.Done()
}
func (fw *FileWorkerPool) consumeResultsWorker() {
	//fw.waitGroup.Add(1)
	for r := range chResults {
		log.Println(r)
	}
	//fw.waitGroup.Done()
}

func (fw *FileWorkerPool) ReadFileContent() {
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
			// Send noOfLines to chLines channel to process. After that clean the lineBuffer.
			log.Println("Send lineBuffer to chLines to be processed.")
			chLines <- lineBuffer
			lineBuffer = []string{}
		}
		line := scanner.Text()
		lineBuffer = append(lineBuffer, line)
	}
	// Check if there is remaining lines in the lineBuffer.
	if len(lineBuffer) != 0 {
		chLines <- lineBuffer
		lineBuffer = nil
	}

	defer file.Close()

}
