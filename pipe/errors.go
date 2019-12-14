package pipe

import (
	"errors"
)

var (
	ErrIgnore        = errors.New("ignore")
	ErrInvalidFormat = errors.New("invalid format")
	ErrInvalidPara   = errors.New("invalid parameter")
)
