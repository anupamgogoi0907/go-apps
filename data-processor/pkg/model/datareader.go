package model

import "sync"

type DataReader struct {
	Path      string
	ChunkPool *sync.Pool
	TextPool  *sync.Pool
	WG        *sync.WaitGroup
}

func NewDataReader(path string) *DataReader {
	chunkPool := sync.Pool{New: func() interface{} {
		buffer := make([]byte, 10*1024)
		return buffer
	}}
	textPool := sync.Pool{New: func() interface{} {
		text := ""
		return text
	}}

	dataReader := &DataReader{
		Path:      path,
		ChunkPool: &chunkPool,
		TextPool:  &textPool,
		WG:        &sync.WaitGroup{},
	}
	return dataReader
}
