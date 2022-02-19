package pipeline

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync"
)

type DataReader struct {
	Path      string
	ChunkPool *sync.Pool
	TextPool  *sync.Pool
	wg        *sync.WaitGroup
}

func InitDataReader(path string) {
	chunkPool := sync.Pool{New: func() interface{} {
		buffer := make([]byte, 10*1024)
		return buffer
	}}
	textPool := sync.Pool{New: func() interface{} {
		text := ""
		return text
	}}

	dataReader := &DataReader{
		Path:      path,
		ChunkPool: &chunkPool,
		TextPool:  &textPool,
		wg:        &sync.WaitGroup{},
	}
	ReadLargeFile(dataReader)
}
func ReadLargeFile(dataReader *DataReader) error {
	// Check for entered file path.
	if dataReader.Path == "" {
		return errors.New("no path found")
	}

	// Check if the file can be opened.
	file, err := os.Open(dataReader.Path)
	if err != nil {
		return err
	}

	// Pools for line
	reader := bufio.NewReader(file)

	// Keep a count of number of chunks.
	nChunks := 0

	// Actual reading.
	flag := true

	for flag {
		chunk := dataReader.ChunkPool.Get().([]byte)
		nBytes, err := reader.Read(chunk)
		if err != nil {
			fmt.Println(err)
			flag = false
			break
		}
		nChunks = nChunks + 1

		dataReader.wg.Add(1)
		go ProcessLine(dataReader, chunk, nBytes, nChunks)
	}
	dataReader.wg.Wait()

	fmt.Println("Total chunks:", nChunks)
	return nil
}

// ProcessLine function is invoked for each chunk concurrently.
func ProcessLine(dataReader *DataReader, chunk []byte, nBytes int, nChunks int) {
	text := dataReader.TextPool.Get().(string)
	text = string(chunk[0:nBytes])

	fmt.Printf("########## Chunk: %d ########## \n%s\n", nChunks, text)

	// Put back chunk and text to the pools
	dataReader.ChunkPool.Put(chunk)
	dataReader.TextPool.Put(text)
	dataReader.wg.Done()
}
