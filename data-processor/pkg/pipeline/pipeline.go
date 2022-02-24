package pipeline

import (
	"errors"
	"fmt"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage"
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

	stageTwo := stage.NewStage("Stage2", func(curStage *stage.Stage) {
		fmt.Println("Processing:", curStage.Name)
		if curStage.Next != nil {
			curStage.Next.Process(curStage.Next)
		}
	}, nil)

	stageOne := stage.NewStage("Stage1", func(curStage *stage.Stage) {
		fmt.Println("Processing:", curStage.Name)
		path := p.Input[0]
		ingest := stage.NewIngest(path, curStage.Data)
		ingest.ReadFileConcurrently()
		if curStage.Next != nil {
			curStage.Next.Process(curStage.Next)
		}
	}, stageTwo)

	stageOne.Process(stageOne)
	return nil
}
