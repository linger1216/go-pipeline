package pipe

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
	_assert(fn != nil, `process func is nil`)
	return &FilterAnonymous{
		name: name,
		fn:   fn,
	}
}
