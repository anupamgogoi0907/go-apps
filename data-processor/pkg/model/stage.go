package model

// Stage is the actual representation of each stage
type Stage struct {
	Id      int
	Name    string
	Process ProcessData
	Next    *Stage
}
type ProcessData func(stage *Stage)

func NewStage(Id int, Name string, Process func(stage *Stage), Next *Stage) *Stage {
	stage := Stage{
		Id:      Id,
		Name:    Name,
		Process: Process,
		Next:    Next,
	}
	return &stage
}

// StageData contains data regarding the input provided to the stages.
type StageData struct {
	FilePath interface{}
	KeyWord  interface{}
}

func NewStageData(FilePath interface{}, KeyWord interface{}) *StageData {
	stageData := StageData{
		FilePath: FilePath,
		KeyWord:  KeyWord,
	}
	return &stageData
}
