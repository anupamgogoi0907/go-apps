package pool

import "github.com/anupamgogoi0907/go-apps/go-docker/pkg/model"

var jobs chan model.Job
var results chan model.Result

func NewPool(sizeBuffer int) {
	jobs = make(chan model.Job, sizeBuffer)
	results = make(chan model.Result, sizeBuffer)
}
