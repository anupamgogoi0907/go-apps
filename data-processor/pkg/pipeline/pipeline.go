package pipeline

import (
	"errors"
	"fmt"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage"
	"sync"
)

type Pipeline struct {
	Input      []string
	wgPipeline *sync.WaitGroup
}

func NewPipeline(Input ...string) (*Pipeline, error) {
	if Input == nil || len(Input) == 0 || len(Input) == 1 {
		return nil, errors.New("no pipeline data provided")
	}
	pipeline := &Pipeline{
		Input:      Input,
		wgPipeline: &sync.WaitGroup{},
	}
	return pipeline, nil
}

func (p *Pipeline) RunPipeline() error {

	stageTwo := stage.NewStage("Stage2", func(curStage *stage.Stage) {
		fmt.Println("Processing:", curStage.Name)
		if curStage.Next != nil {
			curStage.Next.Process(curStage.Next)
		}
	}, nil, p.wgPipeline)

	stageOne := stage.NewStage("Stage1", func(curStage *stage.Stage) {
		fmt.Println("Processing:", curStage.Name)
		path := p.Input[0]
		ingest := stage.NewIngest(string(path))
		ingest.ReadFile()
		if curStage.Next != nil {
			curStage.Next.Process(curStage.Next)
		}
	}, stageTwo, p.wgPipeline)

	stageOne.Process(stageOne)
	p.wgPipeline.Wait()

	return nil
}
