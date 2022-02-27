package pipeline

import (
	"errors"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage/ingest"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage/transform"
	"sync"
)

type Pipeline struct {
	Input []string
}

func NewPipeline(Input ...string) (*Pipeline, error) {
	if Input == nil || len(Input) == 0 {
		return nil, errors.New("no pipeline data provided")
	}
	pipeline := &Pipeline{
		Input: Input,
	}
	return pipeline, nil
}

func (p *Pipeline) RunPipeline() error {
	wg := &sync.WaitGroup{}

	sc := stage.NewStageContextBuilder().StageData(p.Input).Build()
	sb := stage.NewStageBuilder()

	stageProcessor1 := ingest.New().Build()
	s1 := sb.Name("Ingest Data").WG(wg).NoOfWorkers(2).Data(make(chan string)).PrevStage(nil).StageProcessor(stageProcessor1).StageContext(sc).Build()
	s1.RunStage()

	stageProcessor2 := transform.New().Build()
	s2 := sb.Name("Transform").WG(wg).NoOfWorkers(2).Data(nil).PrevStage(s1).StageProcessor(stageProcessor2).StageContext(sc).Build()
	s2.RunStage()

	// Wait for all goroutines that belong to all stages to finish.
	wg.Wait()
	return nil
}
