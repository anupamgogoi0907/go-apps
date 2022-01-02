package pool

import (
	"github.com/anupamgogoi0907/go-apps/go-docker/pkg/model"
	"io/ioutil"
	"log"
)

var jobs chan model.Job
var results chan model.Result

type WorkerPool struct {
	CurDir string
}

func (wp *WorkerPool) configurePool(bufferSize int) {
	jobs = make(chan model.Job, bufferSize)
	results = make(chan model.Result, bufferSize)
}

func (wp *WorkerPool) Run() {
	files := GetFilesOfCurrentDirectory(wp.CurDir)
	bufferSize := len(files)
	wp.configurePool(bufferSize)

}

func GetFilesOfCurrentDirectory(dir string) (arr []string) {
	if dir == "" {
		log.Panicln("Empty string")
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Panicln(err)
	}

	var arrFiles []string
	for _, file := range files {
		arrFiles = append(arrFiles, dir+file.Name())
	}
	return arrFiles
}
