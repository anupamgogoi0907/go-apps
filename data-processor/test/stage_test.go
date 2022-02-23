package test

import (
	"fmt"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/model"
	"testing"
)

func TestStage(t *testing.T) {
	stage2 := model.Stage{
		Id:   2,
		Name: "Read File",
		Process: func(stage *model.Stage) {
			fmt.Println("My stage1 id is:", stage.Id)
		},
		Next: nil,
	}
	stage1 := model.Stage{
		Id:   1,
		Name: "Read File",
		Process: func(stage *model.Stage) {
			fmt.Println("My stage1 id is:", stage.Id)
		},
		Next: &stage2,
	}

	fmt.Printf("Cur:%v, Next:%p\n", stage1, &stage2)

	stage1.Next.Id = 1000
	stage2.Process(&stage2)
}
