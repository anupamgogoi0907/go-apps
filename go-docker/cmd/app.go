package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func ProcessLine(ch chan string) {
	i := 0
	for {
		i = i + 1
		fmt.Println("Receiving line", i)
		line := <-ch
		fmt.Println(line)
	}

}
func ReadFile() {
	ch := make(chan string)
	go ProcessLine(ch)

	fileName, _ := os.Getwd()
	file, err := os.Open(fileName + "/go.sum")
	if err != nil {
		log.Panicln(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		time.Sleep(time.Second * 5)
		ch <- scanner.Text()
	}
}
