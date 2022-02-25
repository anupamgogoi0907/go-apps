package stage

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

const (
	kb = 1024
	mb = kb * 1024
	gb = mb * 1024
)

var (
	chunkSize = 10 * kb
	chunk     = make([]byte, chunkSize)
)

type Ingest struct {
	Path      string
	ChunkPool *sync.Pool
	TextPool  *sync.Pool
	Cur       *Stage
}

func NewStageProcessor(args ...string) *Ingest {
	chunkPool := sync.Pool{New: func() interface{} {
		chunk := chunk
		return chunk
	}}
	textPool := sync.Pool{New: func() interface{} {
		text := ""
		return text
	}}

	in := &Ingest{
		Path:      args[0],
		ChunkPool: &chunkPool,
		TextPool:  &textPool,
	}
	return in
}

func (in *Ingest) RunStageProcessor(cur *Stage) {
	in.Cur = cur
	in.readFileConcurrently()
}
func (in *Ingest) readFileConcurrently() error {
	offset := int64(0)

	// Spawn workers number of goroutines.
	for i := 1; i <= in.Cur.NoOfWorkers; i++ {
		in.Cur.WG.Add(1)
		go in.readFileConcurrentlyRoutine(i, offset)
		offset = offset + int64(chunkSize)
	}
	in.Cur.WG.Wait()
	return nil
}
func (in *Ingest) readFileConcurrentlyRoutine(workerId int, offset int64) error {
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

	in.Cur.WG.Done()
	return nil
}
