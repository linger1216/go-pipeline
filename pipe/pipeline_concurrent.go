package pipe

import (
	"context"
	"fmt"
	"sync"

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

func (c *ConcurrentPipeline) Process(ctx context.Context, req Request) (Response, error) {
	var resp []interface{}
	var err error

	wg := sync.WaitGroup{}
	for i := range c.Filters {
		wg.Add(1)
		go func(pos int) {
			defer wg.Done()
			resp[i], err = c.Filters[i].Process(ctx, req)
			if err != nil && err != common.ErrIgnore {
				panic(fmt.Sprintf("[%s->%s] process error:%s \n", c.name, filter.Name(), err.Error()))
			}
		}(i)
	}
	wg.Wait()

	// go func() {
	// 	for _, filter := range c.Filters {
	// 		if c.debug {
	// 			fmt.Printf("[%s->%s] process\n", c.name, filter.Name())
	// 		}
	// 		resp, err = filter.Process(ctx, req)
	// 		if err != nil && err != common.ErrIgnore {
	// 			panic(fmt.Sprintf("[%s->%s] process error:%s \n", c.name, filter.Name(), err.Error()))
	// 		}
	// 		req = resp
	// 	}
	// }()
	return resp, nil
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
