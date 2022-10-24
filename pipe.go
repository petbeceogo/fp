package fp

type Pipe interface {
	Name() string
	Val() interface{}
}

type pipe struct {
	name string
	val  interface{}
}

func (p *pipe) Name() string {
	return p.name
}

func (p *pipe) Val() interface{} {
	return p.val
}

func NewPipe(name string, val interface{}) Pipe {
	return &pipe{name: name, val: val}
}
