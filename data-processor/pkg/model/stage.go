package model

import "sync"

// Stage is the actual representation of each stage
type Stage struct {
	Name    string
	Process ProcessData
	Next    *Stage
	done    chan bool
	WG      *sync.WaitGroup
}
type ProcessData func(stage *Stage)

func NewStage(Name string, Process func(stage *Stage), Next *Stage, WG *sync.WaitGroup) *Stage {
	stage := Stage{
		Name:    Name,
		Process: Process,
		Next:    Next,
		done:    make(chan bool),
		WG:      WG,
	}
	return &stage
}

// StageData contains data regarding the input provided to the stages.
type StageData struct {
	Input []interface{}
}

func NewStageData(Input ...interface{}) *StageData {
	stageData := StageData{
		Input: Input,
	}
	return &stageData
}
