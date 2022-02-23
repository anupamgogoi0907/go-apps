package test

import (
	"fmt"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/model"
	"testing"
)

func TestStage(t *testing.T) {
	next := model.Stage{
		Id:   2,
		Name: "Read File",
		Process: func(stage *model.Stage) {
			fmt.Println("My stage id is:", stage.Id)
		},
		Next: nil,
	}
	cur := model.Stage{
		Id:   1,
		Name: "Read File",
		Next: &next,
	}

	fmt.Printf("Cur:%v, Next:%p\n", cur, &next)

	cur.Next.Id = 1000
	next.Process(&next)
}
