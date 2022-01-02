package main

import "github.com/anupamgogoi0907/go-apps/go-docker/pkg/pool"

func main() {

	wp := pool.WorkerPool{
		CurDir:      "/Users/agogoi/git/anupamgogoi0907/go-apps/go-docker/test",
		NoOfWorkers: 10,
	}
	wp.Run()
}
