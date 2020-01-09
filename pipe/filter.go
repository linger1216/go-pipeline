package pipe

import (
	"context"
	"github.com/linger1216/go-pipeline/common"
)

type Request interface{}
type Requests []Request

type Response interface{}
type Responses []Response

type ResponsesIterator struct {
	pos int
	arr []Response
}

func ToIterator(arr []Response) *ResponsesIterator {
	return &ResponsesIterator{arr: arr, pos: 0}
}

func (i *ResponsesIterator) Next() (interface{}, error) {
	if i.pos < 0 || i.pos >= len(i.arr) {
		return nil, common.ErrOutOfRange
	}
	ret := i.arr[i.pos]
	i.pos++
	return ret, nil
}

type Process func(ctx context.Context, request Request) (Response, error)

type Filter interface {
	Name() string
	Process(ctx context.Context, req Request) (Response, error)
}
