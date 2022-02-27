package stage

import "sync"

type StageBuilder interface {
	Name(Name string) StageBuilder
	NoOfWorkers(NoOfWorkers int) StageBuilder
	WG(WG *sync.WaitGroup) StageBuilder
	PrevStage(PrevStage *Stage) StageBuilder
	StageProcessor(StageProcessor IStageProcessor) StageBuilder
	StageContext(StageContext *StageContext) StageBuilder
	Build() *Stage
}

type stageBuilder struct {
	name           string
	noOfWorkers    int
	wg             *sync.WaitGroup
	prevStage      *Stage
	stageProcessor IStageProcessor
	stageContext   *StageContext
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

func (sb *stageBuilder) PrevStage(PrevStage *Stage) StageBuilder {
	sb.prevStage = PrevStage
	return sb
}
func (sb *stageBuilder) StageProcessor(StageProcessor IStageProcessor) StageBuilder {
	sb.stageProcessor = StageProcessor
	return sb
}
func (sb *stageBuilder) StageContext(StageContext *StageContext) StageBuilder {
	sb.stageContext = StageContext
	return sb
}
func NewStageBuilder() StageBuilder {
	return &stageBuilder{}
}

func (sb *stageBuilder) Build() *Stage {
	doneWorkers := uint64(0)
	data := make(chan string)
	stage := &Stage{
		Name:           sb.name,
		NoOfWorkers:    sb.noOfWorkers,
		DoneWorkers:    &doneWorkers,
		WG:             sb.wg,
		Data:           data,
		PrevStage:      sb.prevStage,
		StageProcessor: sb.stageProcessor,
		StageContext:   sb.stageContext,
	}
	return stage
}
