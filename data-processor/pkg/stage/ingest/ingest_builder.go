package ingest

import "sync"

type IngestBuilder interface {
	Path(Path string) IngestBuilder
	Build() *Ingest
}

type ingestBuilder struct {
	path string
}

func (ib *ingestBuilder) Path(Path string) IngestBuilder {
	ib.path = Path
	return ib
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
		Path:      ib.path,
		ChunkPool: &chunkPool,
		TextPool:  &textPool,
	}
	return in
}

func New() IngestBuilder {
	ib := &ingestBuilder{}
	return ib
}