package pipe

import (
	"context"
	"testing"
)

func Test_StraightPipeline(t *testing.T) {
	_, err := NewStraight("Straight").
		Append("Straight1", func(ctx context.Context, request Request) (response Response, e error) {
			//fmt.Println("Straight1")
			return nil, nil
		}).Append("Straight2", func(ctx context.Context, request Request) (response Response, e error) {
		//fmt.Println("Straight2")
		return nil, nil
	}).Append("Straight3", func(ctx context.Context, request Request) (response Response, e error) {
		//fmt.Println("Straight3")
		return nil, nil
	}).Append("Straight4", func(ctx context.Context, request Request) (response Response, e error) {
		//fmt.Println("Straight4")
		return nil, nil
	}).Process(nil, nil)

	if err != nil {
		t.Error(err)
	}
}
