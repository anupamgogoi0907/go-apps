package stage

import (
	"sync"
)

type Stage struct {
	Name           string
	NoOfWorkers    int
	DoneWorkers    *uint64
	WG             *sync.WaitGroup
	Data           chan string
	PrevStage      *Stage
	StageProcessor IStageProcessor
	StageContext   *StageContext
}

func (CurStage *Stage) RunStage() {
	CurStage.StageProcessor.RunStageProcessor(CurStage)
}
