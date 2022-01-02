package test

import (
	"fmt"
	"github.com/anupamgogoi0907/go-apps/go-docker/pkg/pool"
	"testing"
)

func TestRun(t *testing.T) {
	wp := pool.WorkerPool{
		CurDir:      "/Users/agogoi/git/anupamgogoi0907/go-apps/go-docker/test",
		NoOfWorkers: 10,
	}
	wp.Run()
}

func TestGetFilesOfCurrentDirectory(t *testing.T) {
	arr := pool.GetFilesOfCurrentDirectory("/Users/agogoi/git/anupamgogoi0907/go-apps/go-docker/test")
	for _, f := range arr {
		fmt.Println(f)
	}
}
