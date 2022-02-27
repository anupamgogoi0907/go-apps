package stage

import "context"

type StageContextBuilder interface {
	Ctx(Ctx *context.Context) StageContextBuilder
	StageData(StageData []string) StageContextBuilder
	Build() *StageContext
}

type stageContextBuilder struct {
	ctx       *context.Context
	stageData []string
}

func (b *stageContextBuilder) Ctx(Ctx *context.Context) StageContextBuilder {
	b.ctx = Ctx
	return b
}
func (b *stageContextBuilder) StageData(StageData []string) StageContextBuilder {
	b.stageData = StageData
	return b
}
func (b *stageContextBuilder) Build() *StageContext {
	sc := &StageContext{
		Ctx:       b.ctx,
		StageData: b.stageData,
	}
	return sc
}
func NewStageContextBuilder() StageContextBuilder {
	sc := &stageContextBuilder{}
	return sc
}
