package pipeline

import (
	"errors"
	"github.com/anupamgogoi0907/go-apps/data-processor/config"
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
	appConfig := config.GetAppConfig()

	stageContext := stage.NewStageContextBuilder().WG(wg).StageData(p.Input).Build()
	sb := stage.NewStageBuilder()

	// Configure stage1.Workers are calculated based on chunk-size.
	stageProcessor1 := ingest.New().ChunkSize(appConfig.Stages[1].Chunksize).Build()
	s1 := sb.Name(appConfig.Stages[1].Name).PrevStage(nil).StageProcessor(stageProcessor1).StageContext(stageContext).Build()
	s1.RunStage()

	// Configure stage2
	stageProcessor2 := transform.New().Build()
	s2 := sb.Name(appConfig.Stages[2].Name).NoOfWorkers(appConfig.Stages[2].Noworkers).PrevStage(s1).StageProcessor(stageProcessor2).StageContext(stageContext).Build()
	s2.RunStage()

	// Wait for all goroutines that belong to all stages to finish.
	wg.Wait()
	return nil
}
