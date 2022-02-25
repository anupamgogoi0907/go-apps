package stage

type IStageProcessor interface {
	RunStageProcessor(cur *Stage)
	//New(args ...string) *IStageProcessor
}
