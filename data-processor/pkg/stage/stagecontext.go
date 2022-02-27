package stage

import "context"

type StageContext struct {
	Ctx       *context.Context
	StageData []string
}
