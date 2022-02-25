package processing

import (
	"bufio"
	"fmt"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage"
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
	CurStage  *stage.Stage
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

func (in *Ingest) RunStageProcessor(CurStage *stage.Stage) {
	in.CurStage = CurStage
	in.readFile()
}
func (in *Ingest) readFile() error {
	offset := int64(0)

	// Spawn workers number of goroutines.
	for i := 1; i <= in.CurStage.NoOfWorkers; i++ {
		in.CurStage.WG.Add(1)
		go in.readFileRoutine(i, offset)
		offset = offset + int64(chunkSize)
	}
	in.CurStage.WG.Wait()
	return nil
}
func (in *Ingest) readFileRoutine(workerId int, offset int64) error {
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

	in.CurStage.WG.Done()
	return nil
}
