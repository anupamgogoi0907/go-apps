package pipeline

import (
	"errors"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage/processing"
	"sync"
)

type Pipeline struct {
	Input []string
}

func NewPipeline(Input ...string) (*Pipeline, error) {
	if Input == nil || len(Input) == 0 || len(Input) == 1 {
		return nil, errors.New("no pipeline data provided")
	}
	pipeline := &Pipeline{
		Input: Input,
	}
	return pipeline, nil
}

func (p *Pipeline) RunPipeline() error {
	wg := &sync.WaitGroup{}

	stageProcessor1 := processing.NewIngestProcessor(p.Input[0])
	s1 := stage.NewStage("Ingest Data", 2, uint64(0), wg, make(chan string), nil, stageProcessor1)
	s1.RunStage()

	stageProcessor2 := processing.NewTransformProcessor("")
	s2 := stage.NewStage("Transform", 2, uint64(0), wg, make(chan string), s1, stageProcessor2)
	s2.RunStage()

	// Wait for all goroutines that belong to all stages to finish.
	wg.Wait()
	return nil
}
