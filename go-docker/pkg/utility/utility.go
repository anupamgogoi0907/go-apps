package utility

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

// GetFilesOfCurrentDirectory finds the files in the provided directory dir
func GetFilesOfCurrentDirectory(dir string) (arr []string) {
	if dir == "" {
		log.Panicln("Empty string")
	}

	dir = AppendLeadingSlash(dir)
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

func CheckLeadingSlash(str string) bool {
	lastChar := str[len(str)-1 : len(str)]
	if lastChar == "/" || lastChar == "\\" {
		return true
	} else {
		return false
	}
}
func AppendLeadingSlash(str string) string {
	if CheckLeadingSlash(str) == false {
		os := runtime.GOOS
		switch os {
		case "windows":
			str = str + "\\"
		case "linux":
			str = str + "/"
		case "darwin":
			str = str + "/"
		}
	}
	return str

}

func ReadFileContent(path string, noOfLines int) {
	if path == "" {
		log.Panicln("No file provided.")
	}
	file, err := os.Open(path)
	if err != nil {
		log.Panicln(err)
	}
	scanner := bufio.NewScanner(file)

	// Count the number of batches of lines sent to chLines
	lineBuffer := []string{}
	for scanner.Scan() {
		if len(lineBuffer) >= noOfLines {
			log.Println("Send lineBuffer to chLines to be processed.")
			lineBuffer = []string{}
		}
		line := scanner.Text()
		lineBuffer = append(lineBuffer, line)

	}
	// Check if there is remaining lines in the lineBuffer.
	if len(lineBuffer) != 0 {
		lineBuffer = nil
	}

	defer file.Close()

}
