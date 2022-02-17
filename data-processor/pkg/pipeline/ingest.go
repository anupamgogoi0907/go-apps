package pipeline

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync"
)

func ReadLargeFile(path string) error {
	// Check for entered file path.
	if path == "" {
		return errors.New("no path found")
	}

	// Check if the file can be opened.
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	// Pools for line
	reader := bufio.NewReader(file)
	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, 500*1024)
		return lines
	}}

	// Keep a count of number of chunks.
	nChunks := 0

	// Actual reading.
	flag := true

	for flag {
		buffer := linesPool.Get().([]byte)
		numBytes, err := reader.Read(buffer)
		if err != nil {
			fmt.Println(err)
			flag = false
			break
		}
		nChunks = nChunks + 1
		line := string(buffer[0:numBytes])
		fmt.Println(line)
	}

	fmt.Println("Total chunks:", nChunks)
	return nil
}
