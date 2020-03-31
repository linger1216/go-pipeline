package async

import (
	"github.com/linger1216/go-pipeline/common"
	"time"
)

type Observable struct {
	name string
	ch   chan interface{}
}

func newObservable(name string, size int) *Observable {
	ret := &Observable{name: name}
	ret.ch = make(chan interface{}, size)
	return ret
}

func _closed(ch <-chan interface{}) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}

func Empty() *Observable {
	ret := newObservable("dummy", 0)
	go func() {
		close(ret.ch)
	}()
	return ret
}

// from

func Interval(done chan struct{}, interval time.Duration) *Observable {
	ret := newObservable("dummy", 0)
	go func(ch chan struct{}) {
		i := 0
	OuterLoop:
		for {
			select {
			case <-ch:
				break OuterLoop
			case <-time.After(interval):
				ret.ch <- i
			}
			i++
		}
		close(ret.ch)
	}(done)
	return ret
}

func Repeat(item interface{}, ntime uint) *Observable {
	if ntime == 0 {
		return Empty()
	}
	ret := newObservable("dummy", 0)
	go func() {
		for i := uint(0); i < ntime; i++ {
			ret.ch <- item
		}
		close(ret.ch)
	}()
	return ret
}

func Concat(ob ...*Observable) *Observable {
	ret := newObservable("dummy", 0)
	go func() {
		for i := range ob {
			for v := range ob[i].ch {
				ret.ch <- v
			}
		}
		close(ret.ch)
	}()
	return ret
}

func Range(start, end int) *Observable {
	ret := newObservable("dummy", 0)
	go func() {
		i := start
		for i < end {
			ret.ch <- i
			i++
		}
		close(ret.ch)
	}()
	return ret
}

func Just(items ...interface{}) *Observable {
	ret := newObservable("dummy", 0)
	go func() {
		for _, item := range items {
			ret.ch <- item
		}
		close(ret.ch)
	}()
	return ret
}

func From(iter common.Iterator) *Observable {
	ret := newObservable("dummy", 0)
	go func() {
		for {
			val, err := iter.Next()
			if err != nil {
				break
			}
			ret.ch <- val
		}
		close(ret.ch)
	}()
	return ret
}

func Create(name string, source func(*Observer)) *Observable {
	ret := newObservable(name, 0)
	closeObservable := func() {
		if !_closed(ret.ch) {
			close(ret.ch)
		}
	}

	nextFn := func(req interface{}) {
		if !_closed(ret.ch) {
			ret.ch <- req
		}
	}

	errorFn := func(err error) {
		if !_closed(ret.ch) {
			ret.ch <- err
		}
		closeObservable()
	}

	completeFn := func() {
		closeObservable()
	}

	emitter := NewObserver(name, OnNextFunction(nextFn), OnErrorFunction(errorFn), OnCompleteFunction(completeFn))

	go func() {
		source(emitter)
	}()
	return ret
}

func (o *Observable) Map(fn MapFunction) *Observable {
	ret := newObservable(o.name, 0)
	go func() {
	Loop:
		for v := range o.ch {
			switch x := v.(type) {
			case error:
				ret.ch <- x
				break Loop
			default:
				ret.ch <- fn(v)
			}
		}
		close(ret.ch)
	}()
	return ret
}

func (o *Observable) Filter(fn FilterFunction) *Observable {
	ret := newObservable(o.name, 0)
	go func() {
		for v := range o.ch {
			if fn(v) {
				ret.ch <- v
			}
		}
		close(ret.ch)
	}()
	return ret
}

func (o *Observable) Take(n uint) *Observable {
	ret := newObservable(o.name, 0)
	go func() {
		count := uint(0)
		for v := range o.ch {
			if count < n {
				ret.ch <- v
			} else {
				break
			}
			count++
		}
		close(ret.ch)
	}()
	return ret
}

func (o *Observable) First() *Observable {
	return o.Take(1)
}

func (o *Observable) Last() *Observable {
	ret := newObservable(o.name, 0)
	go func() {
		var tmp interface{}
		for item := range o.ch {
			tmp = item
		}
		ret.ch <- tmp
		close(ret.ch)
	}()
	return ret
}

func (o *Observable) TakeLast(nth uint) *Observable {
	ret := newObservable(o.name, 0)
	go func() {
		buf := make([]interface{}, nth)
		for item := range o.ch {
			if len(buf) >= int(nth) {
				buf = buf[1:]
			}
			buf = append(buf, item)
		}
		for _, takenItem := range buf {
			ret.ch <- takenItem
		}
		close(ret.ch)
	}()
	return ret
}

func (o *Observable) Distinct(fn SelectFunction) *Observable {
	ret := newObservable(o.name, 0)
	go func() {
		keysets := make(map[interface{}]struct{})
		for item := range o.ch {
			key := fn(item)
			_, ok := keysets[key]
			if !ok {
				ret.ch <- item
			}
			keysets[key] = struct{}{}
		}
		close(ret.ch)
	}()
	return ret
}

func (o *Observable) DistinctUntilChanged(fn SelectFunction) *Observable {
	ret := newObservable(o.name, 0)
	go func() {
		var current interface{}
		for item := range o.ch {
			key := fn(item)
			if current != key {
				ret.ch <- item
				current = key
			}
		}
		close(ret.ch)
	}()
	return ret
}

func (o *Observable) Skip(nth uint) *Observable {
	ret := newObservable(o.name, 0)
	go func() {
		skipCount := 0
		for item := range o.ch {
			if skipCount < int(nth) {
				skipCount += 1
				continue
			}
			ret.ch <- item
		}
		close(ret.ch)
	}()
	return ret
}

func (o *Observable) SkipLast(nth uint) *Observable {
	ret := newObservable(o.name, 0)
	go func() {
		buf := make(chan interface{}, nth)
		for item := range o.ch {
			select {
			case buf <- item:
			default:
				ret.ch <- <-buf
				buf <- item
			}
		}
		close(buf)
		close(ret.ch)
	}()
	return ret
}

func (o *Observable) Scan(fn ScanFunction) *Observable {
	ret := newObservable(o.name, 0)
	go func() {
		var current interface{}
		for item := range o.ch {
			v := fn(current, item)
			ret.ch <- v
			current = v
		}
		close(ret.ch)
	}()
	return ret
}

func (o *Observable) Concat(ob *Observable) *Observable {
	ret := newObservable(o.name, 0)
	go func() {
		for v := range o.ch {
			ret.ch <- v
		}
		for v := range ob.ch {
			ret.ch <- v
		}
		close(ret.ch)
	}()
	return ret
}

func (o *Observable) StartWith(item interface{}) *Observable {
	ret := newObservable(o.name, 0)
	go func() {
		ret.ch <- item
		for v := range o.ch {
			ret.ch <- v
		}
		close(ret.ch)
	}()
	return ret
}

// member functions
func (o *Observable) Subscribe(ob *Observer) *Subscription {
	interrupted := false
	subscription := newSubscription()
	exec := func() {
	Loop:
		for {
			select {
			case v, ok := <-o.ch:
				if !ok {
					// 1. 用户已经发完了 (Create 已经onComplete)
					// 2. Observable 的channel 意外关闭 (不太可能, 属于异常)
					// 这里不调用complete, 留待下面去调用
					break Loop
				}
				switch item := v.(type) {
				case error:
					ob.OnError(item)
					subscription.err = item
					break Loop
				default:
					ob.OnNext(item)
				}
			case <-subscription.done:
				interrupted = true
				break Loop
			}
		}

		if subscription.Error() == nil {
			ob.OnComplete()
		}

		// 关闭subscription
		// 如果是自然结束, 自动调用Unsubscribe, 用户端wait状态结束
		// 如果是非自然结束, 用户手动调用Unsubscribe, 用户端wait状态结束
		if !interrupted {
			subscription.Unsubscribe()
		}
	}

	go func() {
		exec()
	}()

	return subscription
}
