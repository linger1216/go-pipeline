package pipe

import (
	"context"
	"fmt"
	"github.com/linger1216/go-pipeline/common"
)

type FilterAt struct {
	pos int
}

func NewFilterAt(pos int) *FilterAt {
	common.Assert(pos < 0, fmt.Sprintf("%d < 0", pos))
	return &FilterAt{
		pos: pos,
	}
}

func (f *FilterAt) Name() string {
	return ""
}

func (f *FilterAt) Process(ctx context.Context, req Request) (Response, error) {
	requests, ok := req.(Responses)
	if !ok {
		return nil, common.ErrInvalidFormat
	}
	if f.pos < 0 || f.pos >= len(requests) {
		return nil, common.ErrInvalidPara
	}
	return requests[f.pos], nil
}
