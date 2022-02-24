package stage

// Stage is the actual representation of each stage
type Stage struct {
	Name    string
	Process ProcessData
	Next    *Stage
	Data    chan string
}
type ProcessData func(stage *Stage)

func NewStage(Name string, Process func(stage *Stage), Next *Stage) *Stage {
	stage := Stage{
		Name:    Name,
		Process: Process,
		Next:    Next,
		Data:    make(chan string),
	}
	return &stage
}
