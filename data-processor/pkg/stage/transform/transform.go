package transform

import (
	"fmt"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage"
	"os"
	"strconv"
	"sync/atomic"
)

type Transform struct {
	Input      []string
	TargetPath string
	CurStage   *stage.Stage
}

func (t *Transform) RunStageProcessor(curStage *stage.Stage) {
	t.CurStage = curStage

	worker := func(workerId int, t *Transform) {
		flag := true
		var file *os.File
		filePath := t.TargetPath + strconv.Itoa(workerId) + ".log"

		for flag {
			select {
			case text := <-t.CurStage.PrevStage.Data:
				file, _ = os.Create(filePath)
				defer file.Close()
				fmt.Printf("<<<<<<<<<< Stage:%s, Worker:%d\n", t.CurStage.Name, workerId)
				fmt.Println(text)
				file.WriteString(text)
			default:
				c := atomic.LoadUint64(curStage.PrevStage.DoneWorkers)
				if int(c) == curStage.PrevStage.NoOfWorkers {
					flag = false
					fmt.Printf("<<<<<<<<<< DONE:Stage:%s, Worker:%d\n", t.CurStage.Name, workerId)
					t.CurStage.WG.Done()
				}
			}
		}
	}
	for w := 1; w <= t.CurStage.NoOfWorkers; w++ {
		t.CurStage.WG.Add(1)
		go worker(w, t)
	}

}
