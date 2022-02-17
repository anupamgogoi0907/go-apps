package test

import (
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/pipeline"
	"log"
	"testing"
)

func TestReadLargeFile(t *testing.T) {
	err := pipeline.ReadLargeFile("/Users/agogoi/softwares/wso2/AM/wso2am-3.1.0/repository/logs/wso2carbon.log")
	if err != nil {
		log.Println(err)
	}
}
