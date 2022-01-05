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
