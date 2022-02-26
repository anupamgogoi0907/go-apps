package processor

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

type TransformBuilder interface {
	Input(Input ...string) TransformBuilder
	TargetPath(TargetPath string) TransformBuilder
	Build() *Transform
}

type transformBuilder struct {
	input      []string
	targetPath string
	curStage   *stage.Stage
}

func (tb *transformBuilder) Input(Input ...string) TransformBuilder {
	tb.input = Input
	return tb
}
func (tb *transformBuilder) TargetPath(TargetPath string) TransformBuilder {
	tb.targetPath = TargetPath
	return tb
}

func (tb *transformBuilder) Build() *Transform {
	t := &Transform{
		Input:      tb.input,
		TargetPath: tb.targetPath,
	}
	return t
}

func NewTransformBuilder() TransformBuilder {
	tb := &transformBuilder{}
	return tb
}

func (t *Transform) RunStageProcessor(curStage *stage.Stage) {
	t.CurStage = curStage

	worker := func(workerId int, curState *stage.Stage) {
		flag := true
		var file *os.File
		filePath := "/Users/agogoi/Downloads/" + strconv.Itoa(workerId) + ".log"

		for flag {
			select {
			case text := <-curState.PrevStage.Data:
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
		go worker(w, t.CurStage)
	}

}
