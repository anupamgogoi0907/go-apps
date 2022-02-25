package stage

import (
	"context"
	"sync"
)

type Stage struct {
	NoOfWorkers int
	DoneWorkers *uint64
	Ctx         context.Context
	CancelFunc  context.CancelFunc
	WG          *sync.WaitGroup
	Data        chan string
	Error       chan string
}
