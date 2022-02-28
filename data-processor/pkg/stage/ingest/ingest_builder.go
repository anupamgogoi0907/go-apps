package ingest

import "sync"

type IngestBuilder interface {
	ChunkSize(ChunkSize int) IngestBuilder
	Build() *Ingest
}

type ingestBuilder struct {
	chunkSize int
}

func (ib *ingestBuilder) ChunkSize(ChunkSize int) IngestBuilder {
	ib.chunkSize = ChunkSize
	return ib
}
func (ib *ingestBuilder) Build() *Ingest {
	chunkPool := sync.Pool{New: func() interface{} {
		chunk := make([]byte, ib.chunkSize)
		return chunk
	}}
	textPool := sync.Pool{New: func() interface{} {
		text := ""
		return text
	}}
	in := &Ingest{
		ChunkSize: ib.chunkSize,
		ChunkPool: &chunkPool,
		TextPool:  &textPool,
	}
	return in
}

func New() IngestBuilder {
	ib := &ingestBuilder{}
	return ib
}
