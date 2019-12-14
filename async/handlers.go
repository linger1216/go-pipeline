package async

type Handler interface {
	Handle(interface{})
}

type OnNextFunction func(interface{})
type OnErrorFunction func(error)
type OnCompleteFunction func()


func (o OnNextFunction) Handle(req interface{}) {
	switch req.(type) {
	case error:
		return
	}
	o(req)
}


func (o OnErrorFunction) Handle(req interface{}) {
	switch t := req.(type) {
	case error:
		o(t)
	}
}

func (o OnCompleteFunction) Handle(interface{}) {
	o()
}
