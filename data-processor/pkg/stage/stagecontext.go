package stage

import (
	"context"
	"sync"
)

type StageContext struct {
	WG        *sync.WaitGroup
	Ctx       *context.Context
	StageData []string
}
