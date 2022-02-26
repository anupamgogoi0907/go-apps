package main

import (
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/pipeline"
	"os"
)

func main() {
	args := os.Args[1:]
	p, err := pipeline.NewPipeline(args...)
	if err == nil {
		p.RunPipeline()
	}
}
