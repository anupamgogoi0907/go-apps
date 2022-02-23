package pipeline

import (
	"errors"
	"fmt"
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
	wg := &sync.WaitGroup{}

	stageTwo := stage.NewStage("Stage2", func(curStage *stage.Stage) {
		fmt.Println("Processing:", curStage.Name)
		if curStage.Next != nil {
			curStage.Next.Process(curStage.Next)
		}
	}, nil, wg)

	stageOne := stage.NewStage("Stage1", func(curStage *stage.Stage) {
		fmt.Println("Processing:", curStage.Name)
		path := p.Input[0]
		ingest := stage.NewIngest(string(path))
		ingest.ReadLargeFile()
		if curStage.Next != nil {
			// Process Next stage.
			curStage.Next.Process(curStage.Next)
		}
	}, stageTwo, wg)

	stageOne.Process(stageOne)
	wg.Wait()

	return nil
}
