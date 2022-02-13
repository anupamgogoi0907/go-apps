package main

import (
	"fmt"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/pipeline"
)

func main() {
	e := pipeline.ReadLargeFile("")
	fmt.Println(e)
}
