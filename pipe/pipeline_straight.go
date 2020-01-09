package pipe

import (
	"context"
	"fmt"
	"github.com/linger1216/go-pipeline/common"
)

type StraightPipeline struct {
	debug   bool
	name    string
	Filters []Filter
}

func (s *StraightPipeline) Name() string {
	return s.name
}

func (s *StraightPipeline) Process(ctx context.Context, req Request) (Response, error) {
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
			return nil, err
		}
		req = resp
	}
	return resp, err
}

func NewStraight(name string) *StraightPipeline {
	return NewStraightPipeline(true, name)
}

func NewStraightPipeline(debug bool, name string, filters ...Filter) *StraightPipeline {
	return &StraightPipeline{
		debug, name, filters,
	}
}

func (s *StraightPipeline) Append(name string, fn Process) *StraightPipeline {
	common.Assert(fn != nil, `process func is nil`)
	dummyFilter := NewFilterAnonymous(name, fn)
	s.AppendFilter(dummyFilter)
	return s
}

func (s *StraightPipeline) AppendFilter(f Filter) *StraightPipeline {
	common.Assert(f != nil, `filter is nil`)
	s.Filters = append(s.Filters, f)
	return s
}
