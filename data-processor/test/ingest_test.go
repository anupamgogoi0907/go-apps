package test

import (
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/pipeline"
	"testing"
)

func TestReadLargeFile(t *testing.T) {
	pipeline.InitDataReader("/Users/agogoi/softwares/wso2/AM/wso2am-3.1.0/repository/logs/wso2carbon.log")
}
