package pipe

import "fmt"

type StraightPipeline struct {
	debug   bool
	Name    string
	Filters []Filter
}

func (s *StraightPipeline) Process(req Request) (Response, error) {
	var resp interface{}
	var err error
	for _, filter := range s.Filters {
		if s.debug {
			fmt.Printf("[%s-%s] process\n", s.Name, filter.Name())
		}
		resp, err = filter.Process(req)
		if err != nil && err != ErrIgnore {
			if s.debug {
				fmt.Printf("[%s-%s] process error:%s \n", s.Name, filter.Name(), err.Error())
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
	_assert(fn != nil, `process func is nil`)
	dummyFilter := NewFilterAnonymous(name, fn)
	s.AppendFilter(dummyFilter)
	return s
}

func (s *StraightPipeline) AppendFilter(f Filter) *StraightPipeline {
	_assert(f != nil, `filter is nil`)
	s.Filters = append(s.Filters, f)
	return s
}
