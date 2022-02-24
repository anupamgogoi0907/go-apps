package stage

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync"
)

var (
	chunkSize = 10 * 1024
)

type Ingest struct {
	Path      string
	ChunkPool *sync.Pool
	TextPool  *sync.Pool
	WG        *sync.WaitGroup
}

func NewIngest(path string) *Ingest {
	chunkPool := sync.Pool{New: func() interface{} {
		buffer := make([]byte, 10*1024)
		return buffer
	}}
	textPool := sync.Pool{New: func() interface{} {
		text := ""
		return text
	}}

	dataReader := &Ingest{
		Path:      path,
		ChunkPool: &chunkPool,
		TextPool:  &textPool,
		WG:        &sync.WaitGroup{},
	}
	return dataReader
}
func (in *Ingest) ReadLargeFileConcurrently() error {
	// Check if the file can be opened.
	file, err := os.Open(in.Path)
	if err != nil {
		return err
	}

	// Check file size.
	fi, _ := file.Stat()
	fileSize := int(fi.Size())

	// Count number of necessary workers.
	count := fileSize / chunkSize
	if r := fileSize % chunkSize; r != 0 {
		count++
	}
	for i := 1; i <= count; i++ {
		in.WG.Add(1)
	}
	in.WG.Wait()

	return nil
}
func (in *Ingest) ReadLargeFile() error {
	// Check for entered file path.
	if in.Path == "" {
		return errors.New("no path found")
	}

	// Check if the file can be opened.
	file, err := os.Open(in.Path)
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
		chunk := in.ChunkPool.Get().([]byte)
		nBytes, err := reader.Read(chunk)
		if err != nil {
			fmt.Println(err)
			flag = false
			break
		}
		nChunks = nChunks + 1

		in.WG.Add(1)
		go in.processLine(chunk, nBytes, nChunks)
	}
	in.WG.Wait()

	fmt.Println("Total chunks:", nChunks)
	return nil
}

// processLine function is invoked for each chunk concurrently.
func (in *Ingest) processLine(chunk []byte, nBytes int, nChunks int) {
	text := in.TextPool.Get().(string)
	text = string(chunk[0:nBytes])

	fmt.Printf("########## Chunk: %d ########## \n%s\n", nChunks, text)

	// Put back chunk and text to the pools
	in.ChunkPool.Put(chunk)
	in.TextPool.Put(text)
	in.WG.Done()

}
