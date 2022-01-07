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
	NoOfWorkers        int
	waitGroup          *sync.WaitGroup
}

func (fw *FileWorkerPool) Run() {
	fw.waitGroup = &sync.WaitGroup{}
	go fw.ReadFileContent()

	// Allocate NoOfWorkers workers to process file lines.
	for i := 1; i <= fw.NoOfWorkers; i++ {
		fw.waitGroup.Add(1)
		go fw.produceResultsWorker(i)
	}
	fw.waitGroup.Wait()

	log.Println("Executed Workers########")
	close(chResults)
}

// 10 produceResultsWorker are working concurrently to process data sent to chLines channel.
func (fw *FileWorkerPool) produceResultsWorker(workerId int) {
	for l := range chLines {
		// Add processing code.
		log.Println("Worker:", workerId, "is processing.", l)
		//chResults <- []string{"Processed"}
		log.Println("Worker:", workerId, "has processed.")
	}
	fw.waitGroup.Done()
}
func (fw *FileWorkerPool) consumeResultsWorker() {
	for r := range chResults {
		log.Println(r)
	}
}

func (fw *FileWorkerPool) ReadFileContent() {
	fw.waitGroup.Add(1)

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
	close(chLines)
	defer file.Close()
	fw.waitGroup.Done()
}
