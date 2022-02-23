package model

type ProcessData func(stage *Stage)

type Stage struct {
	Id      int
	Name    string
	Process ProcessData
	Next    *Stage
}
