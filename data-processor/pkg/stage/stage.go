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

type StageBuilder interface {
	Name(Name string) StageBuilder
	NoOfWorkers(NoOfWorkers int) StageBuilder
	WG(WG *sync.WaitGroup) StageBuilder
	Data(Data chan string) StageBuilder
	PrevStage(PrevStage *Stage) StageBuilder
	StageProcessor(StageProcessor IStageProcessor) StageBuilder
	Build() *Stage
}

type stageBuilder struct {
	name           string
	noOfWorkers    int
	ctx            context.Context
	cancelFunc     context.CancelFunc
	wg             *sync.WaitGroup
	data           chan string
	error          chan string
	prevStage      *Stage
	stageProcessor IStageProcessor
}

func (sb *stageBuilder) Name(Name string) StageBuilder {
	sb.name = Name
	return sb
}

func (sb *stageBuilder) NoOfWorkers(NoOfWorkers int) StageBuilder {
	sb.noOfWorkers = NoOfWorkers
	return sb
}

func (sb *stageBuilder) WG(WG *sync.WaitGroup) StageBuilder {
	sb.wg = WG
	return sb
}
func (sb *stageBuilder) Data(Data chan string) StageBuilder {
	sb.data = Data
	return sb
}
func (sb *stageBuilder) PrevStage(PrevStage *Stage) StageBuilder {
	sb.prevStage = PrevStage
	return sb
}
func (sb *stageBuilder) StageProcessor(StageProcessor IStageProcessor) StageBuilder {
	sb.stageProcessor = StageProcessor
	return sb
}

func NewStageBuilder() StageBuilder {
	return &stageBuilder{}
}

func (sb *stageBuilder) Build() *Stage {
	doneWorkers := uint64(0)
	stage := &Stage{
		Name:           sb.name,
		NoOfWorkers:    sb.noOfWorkers,
		DoneWorkers:    &doneWorkers,
		Ctx:            sb.ctx,
		CancelFunc:     sb.cancelFunc,
		WG:             sb.wg,
		Data:           sb.data,
		Error:          sb.error,
		PrevStage:      sb.prevStage,
		StageProcessor: sb.stageProcessor,
	}
	return stage
}
