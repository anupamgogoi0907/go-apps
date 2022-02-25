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
	filePath  = "/Users/agogoi/softwares/wso2/AM/wso2am-3.2.0/repository/logs/wso2carbon.log"
	chunkSize = 1024
	chunk     = make([]byte, chunkSize)
)

func TestReadFileConcurrently(t *testing.T) {
	//go func() {
	//	time.Sleep(time.Second)
	//	fmt.Println(<-data)
	//}()
	in := stage.NewStageProcessor("demo.log")
	in.ReadFileConcurrently()
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
	limit := 1
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
	cur = cur + limit
	file.Seek(int64(cur), 0)
	nBytes, _ = reader.Read(chunk)
	chunk = chunk[0:nBytes]
	data = string(chunk)
	assert.Equal(t, "bcde", data)

	// Third read.
	cur = cur + limit
	file.Seek(int64(cur), 0)
	nBytes, _ = reader.Read(chunk)
	chunk = chunk[0:nBytes]
	data = string(chunk)
	assert.Equal(t, "cde", data)
}

func TestComposition(t *testing.T) {
	s1 := &stage.Stage{
		NoOfWorkers: 100,
		DoneWorkers: nil,
		Ctx:         nil,
		CancelFunc:  nil,
		WG:          nil,
		Data:        nil,
		Error:       nil,
	}
	in := stage.Ingest{
		Path:      "",
		ChunkPool: nil,
		TextPool:  nil,
		Cur:       s1,
	}
	assert.Equal(t, 100, in.Cur.NoOfWorkers)
}
