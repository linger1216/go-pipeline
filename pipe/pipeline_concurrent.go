package pipe

import "fmt"

type ConcurrentPipeline struct {
	debug   bool
	Name    string
	Filters []Filter
}

func (c *ConcurrentPipeline) Process(req Request) (Response, error) {
	var resp interface{}
	var err error
	go func() {
		for _, filter := range c.Filters {
			if c.debug {
				fmt.Printf("[%s-%s] process\n", c.Name, filter.Name())
			}
			resp, err = filter.Process(req)
			if err != nil && err != ErrIgnore {
				if c.debug {
					fmt.Printf("[%s-%s] process error:%s \n", c.Name, filter.Name(), err.Error())
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
	_assert(fn != nil, `process func is nil`)
	dummyFilter := NewFilterAnonymous(name, fn)
	c.AppendFilter(dummyFilter)
	return c
}

func (c *ConcurrentPipeline) AppendFilter(f Filter) *ConcurrentPipeline {
	_assert(f != nil, `filter is nil`)
	c.Filters = append(c.Filters, f)
	return c
}
