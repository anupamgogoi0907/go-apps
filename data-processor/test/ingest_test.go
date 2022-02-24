package test

import (
	"bufio"
	"fmt"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var (
	filePath = "/Users/agogoi/softwares/wso2/AM/wso2am-3.2.0/repository/logs/wso2carbon.log"
	chunk    = make([]byte, 1024)
)

func TestReadLargeFile(t *testing.T) {
	d := stage.NewIngest(filePath)
	d.ReadLargeFile()
}

func TestReadLineByLine(t *testing.T) {
	file, _ := os.Open(filePath)
	reader := bufio.NewReader(file)
	bytes := make([]byte, 50*1024)

	fmt.Println("########### 1 ###########")
	nBytes, _ := reader.Read(bytes)
	bytes = bytes[:nBytes]
	text := string(bytes)

	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "2022-02-21 16:07:52") {
			fmt.Println(line)
		}
	}
}

func TestFileOffset(t *testing.T) {
	cur := 0
	file, _ := os.Open("demo.log")
	reader := bufio.NewReader(file)

	// First read.
	file.Seek(int64(cur), 0)
	nBytes, _ := reader.Read(chunk)
	chunk = chunk[0:nBytes]
	data := string(chunk)
	assert.Equal(t, "abcde", data)

	// Second read.
	cur = cur + 1
	file.Seek(int64(cur), 0)
	nBytes, _ = reader.Read(chunk)
	chunk = chunk[0:nBytes]
	data = string(chunk)
	assert.Equal(t, "bcde", data)
}
