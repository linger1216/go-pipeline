package pipe

import (
	"testing"
)

func Test_StraightPipeline(t *testing.T) {
	_, err := NewStraight("Straight").
		Append("Straight1", func(request Request) (response Response, e error) {
			//fmt.Println("Straight1")
			return nil, nil
		}).Append("Straight2", func(request Request) (response Response, e error) {
		//fmt.Println("Straight2")
		return nil, nil
	}).Append("Straight3", func(request Request) (response Response, e error) {
		//fmt.Println("Straight3")
		return nil, nil
	}).Append("Straight4", func(request Request) (response Response, e error) {
		//fmt.Println("Straight4")
		return nil, nil
	}).Process(nil)

	if err != nil {
		t.Error(err)
	}
}
