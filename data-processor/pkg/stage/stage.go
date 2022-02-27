package stage

type Stage struct {
	Name           string
	NoOfWorkers    int
	DoneWorkers    *uint64
	Data           chan string
	PrevStage      *Stage
	StageProcessor IStageProcessor
	StageContext   *StageContext
}

func (CurStage *Stage) RunStage() {
	CurStage.StageProcessor.RunStageProcessor(CurStage)
}
