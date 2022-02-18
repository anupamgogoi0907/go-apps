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
	LinePool  *sync.Pool
}

func InitDataReader(path string) {
	chunkPool := sync.Pool{New: func() interface{} {
		buffer := make([]byte, 50*1024)
		return buffer
	}}
	linePool := sync.Pool{New: func() interface{} {
		lines := ""
		return lines
	}}

	dataReader := &DataReader{
		Path:      path,
		ChunkPool: &chunkPool,
		LinePool:  &linePool,
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
		ProcessLine(dataReader, chunk, nBytes)
	}

	fmt.Println("Total chunks:", nChunks)
	return nil
}

func ProcessLine(dataReader *DataReader, chunk []byte, nBytes int) {
	line := dataReader.LinePool.Get().(string)
	line = string(chunk[0:nBytes])
	fmt.Println(line)

	dataReader.ChunkPool.Put(chunk)
	dataReader.LinePool.Put(line)
}
