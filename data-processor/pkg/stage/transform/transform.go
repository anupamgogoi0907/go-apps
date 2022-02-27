package transform

import (
	"fmt"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage"
	"os"
	"strconv"
	"sync/atomic"
)

type Transform struct {
	SearchStrings []string
	TargetPath    string
	CurStage      *stage.Stage
}

func (t *Transform) RunStageProcessor(curStage *stage.Stage) {
	t.TargetPath = curStage.StageContext.StageData[1]
	t.SearchStrings = curStage.StageContext.StageData[2:]
	t.CurStage = curStage

	t.processData()
}
func (t *Transform) processData() {
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
				c := atomic.LoadUint64(t.CurStage.PrevStage.DoneWorkers)
				if int(c) == t.CurStage.PrevStage.NoOfWorkers {
					flag = false
					fmt.Printf("<<<<<<<<<< DONE:Stage:%s, Worker:%d\n", t.CurStage.Name, workerId)
					t.CurStage.StageContext.WG.Done()
				}
			}
		}
	}
	for w := 1; w <= t.CurStage.NoOfWorkers; w++ {
		t.CurStage.StageContext.WG.Add(1)
		go worker(w, t)
	}
}
