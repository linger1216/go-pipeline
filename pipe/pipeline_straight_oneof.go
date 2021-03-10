package pipe

import (
	"context"
	"fmt"
	"github.com/linger1216/go-pipeline/common"
)

type StraightOneofPipeline struct {
	debug   bool
	name    string
	Filters []Filter
}

func (s *StraightOneofPipeline) Name() string {
	return s.name
}

func (s *StraightOneofPipeline) Process(ctx context.Context, req Request) (Response, error) {
	var resp interface{}
	var err error
	for _, filter := range s.Filters {
		if s.debug {
			fmt.Printf("[%s->%s] process\n", s.name, filter.Name())
		}
		resp, err = filter.Process(ctx, req)
		if err != nil && err != common.ErrIgnore {
			if s.debug {
				fmt.Printf("[%s-%s] process error:%s \n", s.name, filter.Name(), err.Error())
			}
			continue
		}
		break
	}
	return resp, err
}

func NewStraightOneof(name string) *StraightOneofPipeline {
	return NewStraightOneofPipeline(true, name)
}

func NewStraightOneofPipeline(debug bool, name string, filters ...Filter) *StraightOneofPipeline {
	return &StraightOneofPipeline{
		debug, name, filters,
	}
}

func (s *StraightOneofPipeline) Append(name string, fn Process) *StraightOneofPipeline {
	common.Assert(fn != nil, `process func is nil`)
	dummyFilter := NewFilterAnonymous(name, fn)
	s.AppendFilter(dummyFilter)
	return s
}

func (s *StraightOneofPipeline) AppendFilter(f Filter) *StraightOneofPipeline {
	common.Assert(f != nil, `filter is nil`)
	s.Filters = append(s.Filters, f)
	return s
}
