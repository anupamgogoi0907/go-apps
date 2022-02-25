package pipeline

import (
	"errors"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage"
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
	stageProcessor1 := stage.NewStageProcessor(p.Input[0])
	s1 := stage.NewStage(2, uint64(0), &sync.WaitGroup{}, make(chan string), nil, stageProcessor1)
	s1.RunStage()

	return nil
}
