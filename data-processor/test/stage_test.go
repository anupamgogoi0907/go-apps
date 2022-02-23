package test

import (
	"fmt"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStage(t *testing.T) {
	p1 := func(st *model.Stage) {
		fmt.Println("Stage id:", st.Id)
	}
	stage1 := model.NewStage(1, "Stage 1", p1, nil)
	stage1.Process(stage1)
	assert.Equal(t, 1, stage1.Id)
}
func TestStageData(t *testing.T) {
	stageData := model.NewStageData("demopath", "Hello World")
	assert.Equal(t, "demopath", stageData.Input[0])
}
