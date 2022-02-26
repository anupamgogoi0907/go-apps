package transform

import "github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage"

type TransformBuilder interface {
	Input(Input ...string) TransformBuilder
	TargetPath(TargetPath string) TransformBuilder
	Build() *Transform
}

type transformBuilder struct {
	input      []string
	targetPath string
	curStage   *stage.Stage
}

func (tb *transformBuilder) Input(Input ...string) TransformBuilder {
	tb.input = Input
	return tb
}
func (tb *transformBuilder) TargetPath(TargetPath string) TransformBuilder {
	tb.targetPath = TargetPath
	return tb
}

func (tb *transformBuilder) Build() *Transform {
	t := &Transform{
		Input:      tb.input,
		TargetPath: tb.targetPath,
	}
	return t
}

func NewTransformBuilder() TransformBuilder {
	tb := &transformBuilder{}
	return tb
}
