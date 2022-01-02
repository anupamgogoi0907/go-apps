package main

import "github.com/anupamgogoi0907/go-apps/go-docker/pkg/pool"

func main() {

	wp := pool.WorkerPool{
		CurDir:     "",
		BufferSize: 10,
	}
	wp.Run()
}
