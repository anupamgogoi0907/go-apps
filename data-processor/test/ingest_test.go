package test

import (
	"bufio"
	"fmt"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/pipeline"
	"os"
	"testing"
)

func TestReadLargeFile(t *testing.T) {
	pipeline.InitDataReader("/Users/agogoi/softwares/wso2/AM/wso2am-3.1.0/repository/logs/wso2carbon.log")
}

func TestReadLineByLine(t *testing.T) {
	file, _ := os.Open("/Users/agogoi/softwares/wso2/AM/wso2am-3.1.0/repository/logs/wso2carbon.log")
	reader := bufio.NewReader(file)
	bytes := make([]byte, 1*1024)

	fmt.Println("########### 1 ###########")
	nBytes, _ := reader.Read(bytes)
	bytes = bytes[:nBytes]
	text := string(bytes)
	fmt.Println(text)

	fmt.Println("########### 2 ###########")
	nBytes, _ = reader.Read(bytes)
	bytes = bytes[:nBytes]
	text = string(bytes)
	fmt.Println(text)

}
