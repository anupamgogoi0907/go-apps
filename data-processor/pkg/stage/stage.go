package stage

import (
	"context"
	"sync"
)

type Stage struct {
	Name           string
	NoOfWorkers    int
	DoneWorkers    *uint64
	Ctx            context.Context
	CancelFunc     context.CancelFunc
	WG             *sync.WaitGroup
	Data           chan string
	Error          chan string
	PrevStage      *Stage
	StageProcessor IStageProcessor
}

func NewStage(Name string, NoOfWorkers int, DoneWorkers uint64, WG *sync.WaitGroup, Data chan string, PrevStage *Stage, StageProcessor IStageProcessor) *Stage {
	stage := &Stage{
		Name:           Name,
		NoOfWorkers:    NoOfWorkers,
		DoneWorkers:    &DoneWorkers,
		WG:             WG,
		Data:           Data,
		PrevStage:      PrevStage,
		StageProcessor: StageProcessor,
	}
	return stage
}

func (CurStage *Stage) RunStage() {
	CurStage.StageProcessor.RunStageProcessor(CurStage)
}
