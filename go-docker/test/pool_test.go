package test

import (
	"github.com/anupamgogoi0907/go-apps/go-docker/pkg/pool"
	"testing"
)

func TestRun(t *testing.T) {
	wp := pool.WorkerPool{
		CurDir:      "/Users/agogoi/git/anupamgogoi0907/go-apps/go-docker/test/",
		NoOfWorkers: 3,
	}
	wp.Run()
}
