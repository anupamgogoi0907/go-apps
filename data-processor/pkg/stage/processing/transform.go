package processing

import "github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage"

type Transform struct {
	CurStage *stage.Stage
}

func (t *Transform) RunStageProcessor(curStage *stage.Stage) {
	t.CurStage = curStage
}
func NewTransformProcessor(args ...string) *Transform {
	t := &Transform{}
	return t
}
