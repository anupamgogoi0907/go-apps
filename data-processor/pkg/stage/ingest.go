package stage

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var (
	chunkSize = 10 * 1024
	chunk     = make([]byte, chunkSize)
	workers   = 3
)

type Ingest struct {
	Path      string
	ChunkPool *sync.Pool
	TextPool  *sync.Pool
	WG        *sync.WaitGroup
}

func NewIngest(path string, data chan string) *Ingest {
	chunkPool := sync.Pool{New: func() interface{} {
		chunk := chunk
		return chunk
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
func (in *Ingest) ReadFileConcurrently() error {
	offset := int64(0)

	// Spawn workers number of goroutines.
	for i := 1; i <= workers; i++ {
		in.WG.Add(1)
		go in.ReadFileConcurrentlyRoutine(i, offset)
		offset = offset + int64(chunkSize)
	}
	in.WG.Wait()
	return nil
}
func (in *Ingest) ReadFileConcurrentlyRoutine(workerId int, offset int64) error {
	file, _ := os.Open(in.Path)
	defer file.Close()
	file.Seek(offset, 0)

	// Get chunk and text from pools
	chunk := in.ChunkPool.Get().([]byte)
	text := in.TextPool.Get().(string)

	// Get the reader and read file.
	reader := bufio.NewReader(file)
	nBytes, err := reader.Read(chunk)

	if err != nil {
		fmt.Printf("########## Worker:%d, Offset:%d ##########\n%s\n", workerId, offset, string(err.Error()))
	} else {
		text = string(chunk[0:nBytes])
		fmt.Printf("########## Worker:%d, Offset:%d ##########\n%s\n", workerId, offset, text)
		//in.Data <- text
	}

	// Put chunk and text back to the respective pools.
	in.ChunkPool.Put(chunk)
	in.TextPool.Put(text)

	in.WG.Done()
	return nil
}
