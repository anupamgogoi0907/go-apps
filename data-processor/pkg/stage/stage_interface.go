package stage

type IStageProcessor interface {
	RunStageProcessor(cur *Stage)
}
