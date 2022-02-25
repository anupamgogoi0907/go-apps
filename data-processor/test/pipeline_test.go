package test

import (
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/pipeline"
	"testing"
)

func TestStage(t *testing.T) {
	p, _ := pipeline.NewPipeline("/Users/agogoi/Downloads/wso2carbon.log", "")
	p.RunPipeline()
}
