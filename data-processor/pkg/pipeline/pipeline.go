package pipeline

import (
	"errors"
	"fmt"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/model"
	"sync"
)

type Pipeline struct {
	Input []interface{}
}

func NewPipeline(Input ...interface{}) (*Pipeline, error) {
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

	stageTwo := model.NewStage("Stage2", func(curStage *model.Stage) {
		if curStage.Next != nil {
			fmt.Println("Processing:", curStage.Name)
			curStage.Next.Process(curStage.Next)
		}
	}, nil, wg)

	stageOne := model.NewStage("Stage1", func(curStage *model.Stage) {
		if curStage.Next != nil {
			fmt.Println("Processing:", curStage.Name)
			curStage.Next.Process(curStage.Next)
		}
	}, stageTwo, wg)

	stageOne.Process(stageOne)
	wg.Wait()

	return nil
}
