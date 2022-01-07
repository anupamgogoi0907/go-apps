package test

import (
	"github.com/anupamgogoi0907/go-apps/go-docker/pkg/pool"
	"testing"
)

func TestRunFW(t *testing.T) {
	fw := pool.FileWorkerPool{
		FilePath:           "/Users/agogoi/Downloads/demodata.log",
		NoOfLinesToProcess: 1,
		NoOfWorkers:        2,
	}
	fw.Run()
}
