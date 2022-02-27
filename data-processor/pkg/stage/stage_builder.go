package stage

import "sync"

type StageBuilder interface {
	Name(Name string) StageBuilder
	NoOfWorkers(NoOfWorkers int) StageBuilder
	WG(WG *sync.WaitGroup) StageBuilder
	Data(Data chan string) StageBuilder
	PrevStage(PrevStage *Stage) StageBuilder
	StageProcessor(StageProcessor IStageProcessor) StageBuilder
	StageContext(StageContext *StageContext) StageBuilder
	Build() *Stage
}

type stageBuilder struct {
	name           string
	noOfWorkers    int
	wg             *sync.WaitGroup
	data           chan string
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
func (sb *stageBuilder) StageContext(StageContext *StageContext) StageBuilder {
	sb.stageContext = StageContext
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
		WG:             sb.wg,
		Data:           sb.data,
		PrevStage:      sb.prevStage,
		StageProcessor: sb.stageProcessor,
		StageContext:   sb.stageContext,
	}
	return stage
}
