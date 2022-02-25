package stage

import (
	"context"
	"sync"
)

type Stage struct {
	NoOfWorkers    int
	DoneWorkers    *uint64
	Ctx            context.Context
	CancelFunc     context.CancelFunc
	WG             *sync.WaitGroup
	Data           chan string
	Error          chan string
	Prev           *Stage
	StageProcessor IStageProcessor
}

func NewStage(NoOfWorkers int, DoneWorkers uint64, WG *sync.WaitGroup, Data chan string, Prev *Stage, StageProcessor IStageProcessor) *Stage {
	stage := &Stage{
		NoOfWorkers:    NoOfWorkers,
		DoneWorkers:    &DoneWorkers,
		WG:             WG,
		Data:           Data,
		Prev:           Prev,
		StageProcessor: StageProcessor,
	}
	return stage
}

func (cur *Stage) RunStage() {
	cur.StageProcessor.RunStageProcessor(cur)
}
