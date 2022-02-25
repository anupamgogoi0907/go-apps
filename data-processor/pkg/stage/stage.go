package stage

import (
	"context"
	"sync"
)

type Stage struct {
	NoOfWorkers int
	DoneWorkers *uint64
	Ctx         context.Context
	CancelFunc  context.CancelFunc
	WG          *sync.WaitGroup
	Data        chan string
	Error       chan string
	Prev        *Stage
}

func NewStage(NoOfWorkers int, DoneWorkers uint64, WG *sync.WaitGroup, Data chan string, Prev *Stage) *Stage {
	stage := &Stage{
		NoOfWorkers: NoOfWorkers,
		DoneWorkers: &DoneWorkers,
		WG:          WG,
		Data:        Data,
		Prev:        Prev,
	}
	return stage
}

type StageOperation interface {
	Init(prev *Stage)
}

func (cur *Stage) Run(op StageOperation) {
	op.Init(cur.Prev)
}
