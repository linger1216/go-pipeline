package pipe

import (
	"context"
	"testing"
)

func Test_ConcurrentPipeline(t *testing.T) {
	NewConcurrent("Concurrent").
		Append("Concurrent1", func(ctx context.Context, request Request) (response Response, e error) {
			//fmt.Println("Concurrent1")
			return nil, nil
		}).Append("Concurrent2", func(ctx context.Context, request Request) (response Response, e error) {
		//fmt.Println("Concurrent2")
		return nil, nil
	}).Append("Concurrent3", func(ctx context.Context, request Request) (response Response, e error) {
		//fmt.Println("Concurrent3")
		return nil, nil
	}).Append("Concurrent4", func(ctx context.Context, request Request) (response Response, e error) {
		//fmt.Println("Concurrent4")
		return nil, nil
	}).Process(nil, nil)
}
