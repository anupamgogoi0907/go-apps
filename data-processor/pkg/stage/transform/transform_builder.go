package transform

type TransformBuilder interface {
	Build() *Transform
}

type transformBuilder struct {
}

func (tb *transformBuilder) Build() *Transform {
	t := &Transform{}
	return t
}

func New() TransformBuilder {
	tb := &transformBuilder{}
	return tb
}
