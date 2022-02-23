package model

type Stage struct {
	Id      int
	Name    string
	Process func(self *Stage)
	Next    *Stage
}
