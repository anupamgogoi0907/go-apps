package stage

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
	WG        *sync.WaitGroup
}

func NewDataReader(path string) *DataReader {
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
		WG:        &sync.WaitGroup{},
	}
	return dataReader
}

func (dataReader *DataReader) ReadLargeFile() error {
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

		dataReader.WG.Add(1)
		go dataReader.ProcessLine(chunk, nBytes, nChunks)
	}
	dataReader.WG.Wait()

	fmt.Println("Total chunks:", nChunks)
	return nil
}

// ProcessLine function is invoked for each chunk concurrently.
func (dataReader *DataReader) ProcessLine(chunk []byte, nBytes int, nChunks int) {
	text := dataReader.TextPool.Get().(string)
	text = string(chunk[0:nBytes])

	fmt.Printf("########## Chunk: %d ########## \n%s\n", nChunks, text)

	// Put back chunk and text to the pools
	dataReader.ChunkPool.Put(chunk)
	dataReader.TextPool.Put(text)
	dataReader.WG.Done()

}
