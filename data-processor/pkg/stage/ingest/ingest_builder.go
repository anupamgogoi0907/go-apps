package ingest

import "sync"

type IngestBuilder interface {
	Build() *Ingest
}

type ingestBuilder struct {
}

func (ib *ingestBuilder) Build() *Ingest {
	chunkPool := sync.Pool{New: func() interface{} {
		chunk := chunk
		return chunk
	}}
	textPool := sync.Pool{New: func() interface{} {
		text := ""
		return text
	}}
	in := &Ingest{
		ChunkPool: &chunkPool,
		TextPool:  &textPool,
	}
	return in
}

func New() IngestBuilder {
	ib := &ingestBuilder{}
	return ib
}