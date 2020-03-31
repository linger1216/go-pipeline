package async

import (
	"fmt"
	"github.com/linger1216/go-pipeline/common"
	"testing"
	"time"
)

func sameInteger(src []int, dst []int) bool {
	if len(src) != len(dst) {
		return false
	}
	for i := range src {
		if src[i] != dst[i] {
			return false
		}
	}
	return true
}

func sameString(src []string, dst []string) bool {
	if len(src) != len(dst) {
		return false
	}
	for i := range src {
		if src[i] != dst[i] {
			return false
		}
	}
	return true
}

func Test_Create(t *testing.T) {
	src := make([]int, 0)
	sub := Create("test create", func(observer *Observer) {
		for i := 0; i < 10; i++ {
			observer.OnNext(i)
		}
		observer.OnComplete()
	}).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_Empty(t *testing.T) {
	src := make([]int, 0)
	sub := Empty().Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_Interval(t *testing.T) {
	src := make([]int, 0)
	done := make(chan struct{})
	sub := Interval(done, time.Millisecond).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if sameInteger(src, []int{0, 1, 2, 3, 4, 5}) {
			t.Errorf("err result:%v\n", src)
		}
	})))

	go func() {
		time.Sleep(time.Millisecond * 5)
		done <- struct{}{}
	}()

	sub.Wait()
}

func Test_Repeat(t *testing.T) {
	src := make([]string, 0)
	sub := Repeat("xxx", 5).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(string))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameString(src, []string{"xxx", "xxx", "xxx", "xxx", "xxx"}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_Range(t *testing.T) {
	src := make([]int, 0)
	sub := Range(0, 10).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_Concat(t *testing.T) {
	src := make([]int, 0)
	sub := Concat(Range(0, 10), Just(1, 2, 3, 4, 5)).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_Concat1(t *testing.T) {
	src := make([]int, 0)
	sub := Range(0, 10).Concat(Just(1, 2, 3, 4, 5)).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_Just(t *testing.T) {
	src := make([]int, 0)
	sub := Just(1, 2, 3, 4, 5).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{1, 2, 3, 4, 5}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_From(t *testing.T) {
	srcInteger := make([]int, 0)
	sub1 := From(common.NewIntegers([]int{1, 2, 3, 4, 5})).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		srcInteger = append(srcInteger, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(srcInteger, []int{1, 2, 3, 4, 5}) {
			t.Errorf("err result:%v\n", srcInteger)
		}
	})))
	sub1.Wait()

	srcString := make([]string, 0)
	sub2 := From(common.NewStrings([]string{"str1", "str2", "str3", "str4", "str5"})).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		srcString = append(srcString, req.(string))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameString(srcString, []string{"str1", "str2", "str3", "str4", "str5"}) {
			t.Errorf("err result:%v\n", srcString)
		}
	})))
	sub2.Wait()
}

func Test_Subscribe(t *testing.T) {
	src := make([]int, 0)
	sub := Create("test create", func(observer *Observer) {
		for i := 0; i < 10; i++ {
			observer.OnNext(i)
			time.Sleep(time.Millisecond)
		}
		observer.OnComplete()
	}).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if sameInteger(src, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("err result:%v\n", src)
		}
	})))

	go func() {
		time.Sleep(time.Millisecond * 3)
		sub.Unsubscribe()
	}()

	sub.Wait()
}

// todo error
func Test_Error(t *testing.T) {
	src := make([]int, 0)
	sub := Create("test error", func(observer *Observer) {
		for i := 0; i < 10; i++ {
			if i == 5 {
				observer.OnError(fmt.Errorf("lid"))
			} else {
				observer.OnNext(i)
			}
		}
		observer.OnComplete()
	}).Map(func(i interface{}) interface{} {
		return i.(int)
	}).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
		t.Log(req)
	}), OnErrorFunction(func(err error) {
		if !sameInteger(src, []int{0, 1, 2, 3, 4}) {
			t.Errorf("err result:%v\n", src)
		}
	}), OnCompleteFunction(func() {

	})))
	sub.Wait()
}

func Test_Map(t *testing.T) {
	src := make([]int, 0)
	sub := Create("test error", func(observer *Observer) {
		for i := 0; i < 10; i++ {
			observer.OnNext(i)
		}
		observer.OnComplete()
	}).Map(func(i interface{}) interface{} {
		return i.(int) + 100
	}).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{100, 101, 102, 103, 104, 105, 106, 107, 108, 109}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_Filter(t *testing.T) {
	src := make([]int, 0)
	sub := Create("test error", func(observer *Observer) {
		for i := 0; i < 10; i++ {
			observer.OnNext(i)
		}
		observer.OnComplete()
	}).Filter(func(i interface{}) bool {
		return i.(int) == 5
	}).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{5}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_Take(t *testing.T) {
	src := make([]int, 0)
	sub := Create("test error", func(observer *Observer) {
		for i := 0; i < 10; i++ {
			observer.OnNext(i)
		}
		observer.OnComplete()
	}).Take(3).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{0, 1, 2}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_First(t *testing.T) {
	src := make([]int, 0)
	sub := Create("test error", func(observer *Observer) {
		for i := 0; i < 10; i++ {
			observer.OnNext(i)
		}
		observer.OnComplete()
	}).First().Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{0}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_Last(t *testing.T) {
	src := make([]int, 0)
	sub := Create("test error", func(observer *Observer) {
		for i := 0; i < 10; i++ {
			observer.OnNext(i)
		}
		observer.OnComplete()
	}).Last().Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{9}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_TakeLast(t *testing.T) {
	src := make([]int, 0)
	sub := Create("test error", func(observer *Observer) {
		for i := 0; i < 10; i++ {
			observer.OnNext(i)
		}
		observer.OnComplete()
	}).TakeLast(3).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{7, 8, 9}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_Distinct(t *testing.T) {
	src := make([]int, 0)
	sub := Create("test error", func(observer *Observer) {
		for i := 0; i < 10; i++ {
			observer.OnNext(i)
		}

		for i := 0; i < 10; i++ {
			observer.OnNext(i)
		}
		observer.OnComplete()
	}).Distinct(func(i interface{}) interface{} {
		return i
	}).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_DistinctUntilChanged(t *testing.T) {
	src := make([]int, 0)
	sub := Just(1, 1, 2, 3, 4, 4, 4, 5).DistinctUntilChanged(func(i interface{}) interface{} {
		return i
	}).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{1, 2, 3, 4, 5}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_Skip(t *testing.T) {
	src := make([]int, 0)
	sub := Create("test error", func(observer *Observer) {
		for i := 0; i < 10; i++ {
			observer.OnNext(i)
		}
		observer.OnComplete()
	}).Skip(3).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_SkipLast(t *testing.T) {
	src := make([]int, 0)
	sub := Create("test error", func(observer *Observer) {
		for i := 0; i < 10; i++ {
			observer.OnNext(i)
		}
		observer.OnComplete()
	}).SkipLast(7).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{0, 1, 2}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_Scan(t *testing.T) {
	src := make([]int, 0)
	sub := Create("test error", func(observer *Observer) {
		for i := 0; i < 10; i++ {
			observer.OnNext(i)
		}
		observer.OnComplete()
	}).Scan(func(i1 interface{}, i2 interface{}) interface{} {
		v1, _ := i1.(int)
		v2, _ := i2.(int)
		return v1 + v2
	}).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{0, 1, 3, 6, 10, 15, 21, 28, 36, 45}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_Throw(t *testing.T) {
	src := make([]int, 0)
	sub := Create("test error", func(observer *Observer) {
		for i := 0; i < 10; i++ {
			observer.OnNext(i)
		}
		observer.OnComplete()
	}).Map(func(i interface{}) interface{} {
		if i.(int) == 5 {
			return fmt.Errorf("lid 5")
		}
		return i.(int) + 100
	}).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {

	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{100, 101, 102, 103, 104}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}

func Test_StartWith(t *testing.T) {
	src := make([]int, 0)
	sub := Create("test error", func(observer *Observer) {
		for i := 0; i < 10; i++ {
			observer.OnNext(i)
		}
		observer.OnComplete()
	}).StartWith(999).Subscribe(NewObserver("", OnNextFunction(func(req interface{}) {
		src = append(src, req.(int))
	}), OnErrorFunction(func(err error) {
		t.Errorf("err:%s", err.Error())
	}), OnCompleteFunction(func() {
		if !sameInteger(src, []int{999, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("err result:%v\n", src)
		}
	})))
	sub.Wait()
}
