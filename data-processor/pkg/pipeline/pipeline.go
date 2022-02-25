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
	s1 := stage.NewStage(2, uint64(0), &sync.WaitGroup{}, make(chan string), nil)
	ingest := stage.NewIngest(p.Input[0], s1)
	s1.Run(ingest)

	return nil
}
