package stage

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var (
	chunkSize = 2
	chunk     = make([]byte, chunkSize)
)

type Ingest struct {
	Path      string
	ChunkPool *sync.Pool
	TextPool  *sync.Pool
	WG        *sync.WaitGroup
	Finished  bool
	data      chan string
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
		Finished:  false,
		data:      data,
	}
	return dataReader
}
func (in *Ingest) ReadFileConcurrently() error {
	// Check if the file can be opened.
	file, err := os.Open(in.Path)
	defer file.Close()
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
	fmt.Println("Total workers:", count)
	// Reader.
	offset := int64(0)
	reader := bufio.NewReader(file)

	// Spawn count number of goroutines.
	for i := 1; i <= count; i++ {
		in.WG.Add(1)
		go in.ReadFileConcurrentlyRoutine(i, offset, file, reader)
		offset = offset + int64(chunkSize)
	}
	in.WG.Wait()
	return nil
}
func (in *Ingest) ReadFileConcurrentlyRoutine(workerId int, offset int64, file *os.File, reader *bufio.Reader) error {
	file.Seek(offset, 0)
	chunk := in.ChunkPool.Get().([]byte)
	nBytes, err := reader.Read(chunk)

	if err != nil {
		fmt.Printf("########## Worker:%d, Offset:%d ##########\n%s\n", workerId, offset, string(err.Error()))
	} else {
		text := in.TextPool.Get().(string)
		text = string(chunk[0:nBytes])
		fmt.Printf("########## Worker:%d, Offset:%d ##########\n%s\n", workerId, offset, text)
		in.TextPool.Put(text)
	}

	in.ChunkPool.Put(chunk)
	in.WG.Done()
	return nil
}
