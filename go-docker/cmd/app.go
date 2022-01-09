package main

import "github.com/anupamgogoi0907/go-apps/go-docker/pkg/pool"

func main() {
	fw := pool.FileWorkerPool{
		FilePath:           "/Users/agogoi/Downloads/demodata.log",
		NoOfLinesToProcess: 3,
		NoOfWorkers:        2,
		TargetLocation:     "/Users/agogoi/Downloads/demo.log",
	}
	fw.Run()
}
