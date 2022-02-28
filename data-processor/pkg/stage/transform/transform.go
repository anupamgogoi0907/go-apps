package transform

import (
	"bufio"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/stage"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/utility"
	"log"
	"os"
	"strings"
	"sync/atomic"
)

type Transform struct {
	SearchStrings []string
	TargetPath    string
	CurStage      *stage.Stage
}

func (t *Transform) RunStageProcessor(curStage *stage.Stage) {
	t.TargetPath = curStage.StageContext.StageData[1]
	t.SearchStrings = curStage.StageContext.StageData[2:]
	t.CurStage = curStage

	t.processData()
}
func (t *Transform) processData() {
	worker := func(workerId int, mapKeyFile map[string]*os.File, t *Transform) {
		flag := true
		for flag {
			select {
			case text := <-t.CurStage.PrevStage.Data:
				log.Printf("<<<<<<<<<< Stage:%s, Worker:%d\n", t.CurStage.Name, workerId)
				//log.Println(text)
				t.searchStringInText(text, mapKeyFile)
			default:
				c := atomic.LoadUint64(t.CurStage.PrevStage.DoneWorkers)
				if int(c) == t.CurStage.PrevStage.NoOfWorkers {
					flag = false
					log.Printf("<<<<<<<<<< DONE:Stage:%s, Worker:%d\n", t.CurStage.Name, workerId)
					t.CurStage.StageContext.WG.Done()

					// Close files.
					if mapKeyFile != nil {
						for _, v := range mapKeyFile {
							v.Close()
						}
					}
				}
			}
		}
	}

	mapKeyFile := t.createFilePerSearchKey()
	for w := 1; w <= t.CurStage.NoOfWorkers; w++ {
		t.CurStage.StageContext.WG.Add(1)
		go worker(w, mapKeyFile, t)
	}
}
func (t *Transform) createFilePerSearchKey() map[string]*os.File {
	err := utility.CreateDir(t.TargetPath)
	if err != nil {
		return nil
	}
	mapKeyFile := make(map[string]*os.File)
	for i := 0; i < len(t.SearchStrings); i++ {
		s := t.SearchStrings[i]
		filePath := t.TargetPath + s + ".log"

		// Create file.
		f, _ := os.Create(filePath)
		log.Println("Created file:", filePath)
		mapKeyFile[s] = f
	}
	return mapKeyFile
}

func (t *Transform) searchStringInText(text string, m map[string]*os.File) {
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		for k, v := range m {
			if strings.Contains(line, k) {
				log.Println("$$$$$$$$$$$ Writing to file:", k)
				v.WriteString(line + "\n")
			}
		}
	}
}
