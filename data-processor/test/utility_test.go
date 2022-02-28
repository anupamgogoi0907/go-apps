package test

import (
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/utility"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateDir(t *testing.T) {
	utility.CreateDir("/Users/agogoi/Downloads/aa")
}

func TestSplitSize(t *testing.T) {
	s := "10MB"
	v, u := utility.SplitSize(s)
	assert.Equal(t, 10, v)
	assert.Equal(t, utility.MB, u)
}
