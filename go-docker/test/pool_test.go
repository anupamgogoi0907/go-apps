package test

import (
	"fmt"
	"github.com/anupamgogoi0907/go-apps/go-docker/pkg/pool"
	"github.com/anupamgogoi0907/go-apps/go-docker/pkg/utility"
	"testing"
)

func TestRun(t *testing.T) {
	wp := pool.WorkerPool{
		CurDir:      "/Users/agogoi/git/anupamgogoi0907/go-apps/go-docker/test/",
		NoOfWorkers: 3,
	}
	wp.Run()
}

func TestGetFilesOfCurrentDirectory(t *testing.T) {
	arr := utility.GetFilesOfCurrentDirectory("/Users/agogoi/git/anupamgogoi0907/go-apps/go-docker/test")
	for _, f := range arr {
		fmt.Println(f)
	}
}
