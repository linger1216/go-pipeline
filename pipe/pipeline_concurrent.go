package pipe

import (
	"fmt"
	"github.com/linger1216/go-pipeline/common"
)

type ConcurrentPipeline struct {
	debug   bool
	name    string
	Filters []Filter
}

func (c *ConcurrentPipeline) Name() string {
	return c.name
}

func (c *ConcurrentPipeline) Process(req Request) (Response, error) {
	var resp interface{}
	var err error
	go func() {
		for _, filter := range c.Filters {
			if c.debug {
				fmt.Printf("[%s->%s] process\n", c.name, filter.Name())
			}
			resp, err = filter.Process(req)
			if err != nil && err != common.ErrIgnore {
				if c.debug {
					fmt.Printf("[%s->%s] process error:%s \n", c.name, filter.Name(), err.Error())
				}
				panic(err)
			}
			req = resp
		}
	}()
	return nil, nil
}

func NewConcurrent(name string) *ConcurrentPipeline {
	return NewConcurrentPipeline(true, name)
}

func NewConcurrentPipeline(debug bool, name string, filters ...Filter) *ConcurrentPipeline {
	return &ConcurrentPipeline{
		debug, name, filters,
	}
}

func (c *ConcurrentPipeline) Append(name string, fn Process) *ConcurrentPipeline {
	common.Assert(fn != nil, `process func is nil`)
	dummyFilter := NewFilterAnonymous(name, fn)
	c.AppendFilter(dummyFilter)
	return c
}

func (c *ConcurrentPipeline) AppendFilter(f Filter) *ConcurrentPipeline {
	common.Assert(f != nil, `filter is nil`)
	c.Filters = append(c.Filters, f)
	return c
}
