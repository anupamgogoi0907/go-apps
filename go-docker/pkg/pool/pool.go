package pool

import (
	"github.com/anupamgogoi0907/go-apps/go-docker/pkg/model"
	"io/ioutil"
	"log"
)

var jobs chan model.Job
var results chan model.Result

type WorkerPool struct {
	CurDir     string
	BufferSize int
}

func (wp *WorkerPool) configurePool() {
	jobs = make(chan model.Job, wp.BufferSize)
	results = make(chan model.Result, wp.BufferSize)
}

func (wp *WorkerPool) Run() {
	wp.configurePool()
	files := GetFilesOfCurrentDirectory(wp.CurDir)
	log.Println(len(files))

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
