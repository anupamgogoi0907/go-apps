package model

// Stage data structure
type ProcessData func(stage *Stage)
type Stage struct {
	Id      int
	Name    string
	Process ProcessData
	Next    *Stage
}

func NewStage(Id int, Name string, Process func(stage *Stage), Next *Stage) *Stage {
	stage := Stage{
		Id:      Id,
		Name:    Name,
		Process: Process,
		Next:    Next,
	}
	return &stage
}

// StageData
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
