package pipeline

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadLargeFile(path string) error {
	if path == "" {
		return errors.New("no path found")
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 4*1024)

	flag := true
	for flag {
		n, err := reader.Read(buffer)
		if err != nil {
			fmt.Println(err)
			flag = false
		}
		line := string(buffer[0:n])
		fmt.Println(line)
	}
	return nil
}
