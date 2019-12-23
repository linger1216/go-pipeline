package pipe

import "github.com/linger1216/go-pipeline/common"

type FilterAnonymous struct {
	name string
	fn   Process
}

func (f *FilterAnonymous) Name() string {
	return f.name
}

func (f *FilterAnonymous) Process(req Request) (Response, error) {
	return f.fn(req)
}

func NewFilterAnonymous(name string, fn Process) *FilterAnonymous {
	common.Assert(fn != nil, `process func is nil`)
	return &FilterAnonymous{
		name: name,
		fn:   fn,
	}
}
