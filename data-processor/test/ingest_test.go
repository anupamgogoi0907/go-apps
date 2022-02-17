package test

import (
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/pipeline"
	"log"
	"testing"
)

func TestReadLargeFile(t *testing.T) {
	err := pipeline.ReadLargeFile("/Users/agogoi/Downloads/test.txt")
	if err != nil {
		log.Println(err)
	}
}
