package async

type Observer struct {
	name     string
	next     OnNextFunction
	err      OnErrorFunction
	complete OnCompleteFunction
}

func NewObserver(name string, handlers ...Handler) *Observer {
	ret := &Observer{name: name}
	for i := range handlers {
		switch fn := handlers[i].(type) {
		case OnNextFunction:
			ret.next = fn
		case OnErrorFunction:
			ret.err = fn
		case OnCompleteFunction:
			ret.complete = fn
		}
	}
	return ret
}

func (o *Observer) OnNext(req interface{}) {
	switch req.(type) {
	case error:
		return
	}
	o.next(req)
}

func (o *Observer) OnError(err error) {
	if o.err != nil {
		o.err(err)
	}
}

func (o *Observer) OnComplete() {
	if o.complete != nil {
		o.complete()
	}
}
