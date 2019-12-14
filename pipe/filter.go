package pipe

type Request interface{}
type Requests []Request

type Response interface{}
type Responses []Response

type Process func(request Request) (Response, error)

type Filter interface {
	Name() string
	Process(req Request) (Response, error)
}
