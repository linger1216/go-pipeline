package pipe

import (
	"fmt"
	"github.com/linger1216/go-pipeline/common"
	"sync"
)

type ParallelPipeline struct {
	debug   bool
	name    string
	Filters []Filter
}

func (p *ParallelPipeline) Name() string {
	return p.name
}

func (p *ParallelPipeline) Process(req Request) (Response, error) {
	wg := sync.WaitGroup{}
	ret := make(Responses, len(p.Filters))
	for i := range p.Filters {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if p.debug {
				fmt.Printf("[%s->%s] process\n", p.name, p.Filters[i].Name())
			}
			if resp, err := p.Filters[i].Process(req); err == nil {
				ret[i] = resp
			} else {
				if p.debug {
					fmt.Printf("[%s->%s] process error:%s \n", p.name, p.Filters[i].Name(), err.Error())
				}
			}
		}(i)
	}
	wg.Wait()
	return ret, nil
}

func NewParallel(name string) *ParallelPipeline {
	return NewParallelPipeline(true, name)
}

func NewParallelPipeline(debug bool, name string, filters ...Filter) *ParallelPipeline {
	return &ParallelPipeline{
		debug, name, filters,
	}
}

func (p *ParallelPipeline) Append(name string, fn Process) *ParallelPipeline {
	common.Assert(fn != nil, `process func is nil`)
	dummyFilter := NewFilterAnonymous(name, fn)
	p.AppendFilter(dummyFilter)
	return p
}

func (p *ParallelPipeline) AppendFilter(f Filter) *ParallelPipeline {
	common.Assert(f != nil, `filter is nil`)
	p.Filters = append(p.Filters, f)
	return p
}
