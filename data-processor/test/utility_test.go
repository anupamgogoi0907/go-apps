package test

import (
	"fmt"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/utility"
	"testing"
)

func TestCreateDir(t *testing.T) {
	utility.CreateDir("/Users/agogoi/Downloads/aa")
}

func Test1(t *testing.T) {
	b := Hello()
	fmt.Println(b)
}

func Hello() (a int) {
	a = 10
	return a
}
