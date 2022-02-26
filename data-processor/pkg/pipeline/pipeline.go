package pipeline

import (
	"errors"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage/processor"
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
	builder := stage.NewStageBuilder()

	stageProcessor1 := processor.NewIngestProcessor(p.Input[0])
	s1 := builder.Name("Ingest Data").WG(wg).NoOfWorkers(2).Data(make(chan string)).PrevStage(nil).StageProcessor(stageProcessor1).Build()
	s1.RunStage()

	stageProcessor2 := processor.NewTransformBuilder().TargetPath(p.Input[1]).Input(p.Input[2:]...).Build()
	s2 := builder.Name("Transform").WG(wg).NoOfWorkers(2).Data(nil).PrevStage(s1).StageProcessor(stageProcessor2).Build()
	s2.RunStage()

	// Wait for all goroutines that belong to all stages to finish.
	wg.Wait()
	return nil
}
