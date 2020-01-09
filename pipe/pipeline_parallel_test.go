package pipe

import (
	"context"
	"testing"
)

func Test_ParallelPipeline(t *testing.T) {
	NewParallel("Parallel").
		Append("Parallel1", func(ctx context.Context, request Request) (response Response, e error) {
			//fmt.Println("Parallel1")
			return nil, nil
		}).Append("Parallel2", func(ctx context.Context, request Request) (response Response, e error) {
		//fmt.Println("Parallel2")
		return nil, nil
	}).Append("Parallel3", func(ctx context.Context, request Request) (response Response, e error) {
		//fmt.Println("Parallel3")
		return nil, nil
	}).Append("Parallel4", func(ctx context.Context, request Request) (response Response, e error) {
		//fmt.Println("Parallel4")
		return nil, nil
	}).Process(nil, nil)
}
