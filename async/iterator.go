package async

// if have other object, need implement this interface
type Iterator interface {
	Next() (interface{}, error)
}

type Strings struct {
	pos int
	arr []string
}

func NewStrings(arr []string) *Strings {
	return &Strings{arr: arr, pos: 0}
}

func (s *Strings) Next() (interface{}, error) {
	if s.pos < 0 || s.pos >= len(s.arr) {
		return nil, ErrOutOfRange
	}
	ret := s.arr[s.pos]
	s.pos++
	return ret, nil
}

type Integers struct {
	pos int
	arr []int
}

func NewIntegers(arr []int) *Integers {
	return &Integers{arr: arr, pos: 0}
}

func (i *Integers) Next() (interface{}, error) {
	if i.pos < 0 || i.pos >= len(i.arr) {
		return nil, ErrOutOfRange
	}
	ret := i.arr[i.pos]
	i.pos++
	return ret, nil
}
