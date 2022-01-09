package test

import (
	"fmt"
	"github.com/anupamgogoi0907/go-apps/go-docker/pkg/utility"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFilesOfCurrentDirectory(t *testing.T) {
	arr := utility.GetFilesOfCurrentDirectory("/Users/agogoi/git/anupamgogoi0907/go-apps/go-docker/test")
	for _, f := range arr {
		fmt.Println(f)
	}
	assert.NotEqual(t, 0, len(arr))
}
func TestCheckLeadingSlash(t *testing.T) {
	dir := "/Users/agogoi/git/anupamgogoi0907/go-apps/go-docker/test"
	res := utility.CheckLeadingSlash(dir)
	assert.Equal(t, false, res)
}
func TestAppendLeadingSlash(t *testing.T) {
	str := "/Users/agogoi/git/anupamgogoi0907/go-apps/go-docker/test"
	str = utility.AppendLeadingSlash(str)
	assert.Equal(t, "/Users/agogoi/git/anupamgogoi0907/go-apps/go-docker/test/", str)
}

func TestReadFileContents(t *testing.T) {
	utility.ReadFileContent("/Users/agogoi/Downloads/demodata.log", 1)
}

func TestFindWord(t *testing.T) {
	line := "TID: [-1234] [] [2022-01-05 15:14:12,163] ERROR {org.wso2.carbon.apimgt.gateway.handlers.security.APIAuthenticationHandler} -  API authentication failed with error 900901 {org.wso2.carbon.apimgt.gateway.handlers.security.APIAuthenticationHandler}"
	word := "ERROR"
	found := utility.FindWord(line, word)
	assert.Equal(t, true, found)
}

func TestCreateOrOpenFile(t *testing.T) {
	path := "demo.log"
	file := utility.CreateOrOpenFile(path)
	defer file.Close()
	assert.NotEmpty(t, file)
}

func TestAppendDataToFile(t *testing.T) {
	path := "demo.log"
	file := utility.CreateOrOpenFile(path)
	defer file.Close()
	assert.NotEmpty(t, file)

	utility.AppendDataToFile(file, "Hello World")
}