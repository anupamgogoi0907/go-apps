package stage

type IStageProcessor interface {
	RunStageProcessor(CurStage *Stage)
}
