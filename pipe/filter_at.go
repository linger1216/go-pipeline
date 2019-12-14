package pipe

import "fmt"

type FilterAt struct {
	pos int
}

func NewFilterAt(pos int) *FilterAt {
	_assert(pos < 0, fmt.Sprintf("%d < 0", pos))
	return &FilterAt{
		pos: pos,
	}
}

func (f *FilterAt) Process(req Request) (Response, error) {
	requests, ok := req.(Responses)
	if !ok {
		return nil, ErrInvalidFormat
	}
	if f.pos < 0 || f.pos >= len(requests) {
		return nil, ErrInvalidPara
	}
	return requests[f.pos], nil
}
