package stage

import (
	"context"
	"sync"
)

type StageContextBuilder interface {
	WG(WG *sync.WaitGroup) StageContextBuilder
	Ctx(Ctx *context.Context) StageContextBuilder
	StageData(StageData []string) StageContextBuilder
	Build() *StageContext
}

type stageContextBuilder struct {
	wg        *sync.WaitGroup
	ctx       *context.Context
	stageData []string
}

func (b *stageContextBuilder) WG(WG *sync.WaitGroup) StageContextBuilder {
	b.wg = WG
	return b
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
		WG:        b.wg,
		Ctx:       b.ctx,
		StageData: b.stageData,
	}
	return sc
}
func NewStageContextBuilder() StageContextBuilder {
	sc := &stageContextBuilder{}
	return sc
}
