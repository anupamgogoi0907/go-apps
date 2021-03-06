package ingest

import (
	"bufio"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage"
	"log"
	"os"
	"sync"
	"sync/atomic"
)

type Ingest struct {
	Path      string
	ChunkSize int
	ChunkPool *sync.Pool
	TextPool  *sync.Pool
	CurStage  *stage.Stage
}

func (in *Ingest) RunStageProcessor(CurStage *stage.Stage) {
	in.CurStage = CurStage
	in.Path = string(CurStage.StageContext.StageData[0])
	in.readFile()
}
func (in *Ingest) readFile() error {
	offset := int64(0)

	// Calculate no of necessary goroutines
	in.CurStage.NoOfWorkers = in.getNoOfWorkers()

	// Spawn workers number of goroutines.
	for i := 1; i <= in.CurStage.NoOfWorkers; i++ {
		in.CurStage.StageContext.WG.Add(1)
		go in.readFileRoutine(i, offset)
		offset = offset + int64(in.ChunkSize)
	}
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
		log.Printf(">>>>>>>>>> Stage:%s, Worker:%d, Offset:%d\n", in.CurStage.Name, workerId, offset)
		log.Println(err)
	} else {
		text = string(chunk[0:nBytes])
		log.Printf(">>>>>>>>>> Stage:%s, Worker:%d, Offset:%d\n", in.CurStage.Name, workerId, offset)
		in.CurStage.Data <- text
	}

	// Put chunk and text back to the respective pools.
	in.ChunkPool.Put(chunk)
	in.TextPool.Put(text)

	in.CurStage.StageContext.WG.Done()
	atomic.AddUint64(in.CurStage.DoneWorkers, 1)
	return nil
}

func (in *Ingest) getNoOfWorkers() int {
	file, _ := os.Open(in.Path)
	defer file.Close()
	fi, _ := file.Stat()
	fileSize := int(fi.Size())

	NoOfWorkers := fileSize / in.ChunkSize
	if r := fileSize % in.ChunkSize; r != 0 {
		NoOfWorkers++
	}
	log.Printf("########## Stage:%s,File Size:%d,Total Workers:%d\n", in.CurStage.Name, fileSize, NoOfWorkers)
	return NoOfWorkers
}
