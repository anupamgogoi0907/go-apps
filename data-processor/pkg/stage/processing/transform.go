package processing

import (
	"fmt"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage"
	"sync/atomic"
)

type Transform struct {
	CurStage *stage.Stage
}

func NewTransformProcessor(args ...string) *Transform {
	t := &Transform{}
	return t
}
func (t *Transform) RunStageProcessor(curStage *stage.Stage) {
	t.CurStage = curStage

	worker := func(workerId int, curState *stage.Stage) {
		flag := true
		for flag {
			select {
			case text := <-curState.PrevStage.Data:
				fmt.Printf("<<<<<<<<<< Stage:%s, Worker:%d\n", t.CurStage.Name, workerId)
				fmt.Println(text)
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
		go worker(w, t.CurStage)
	}

}
