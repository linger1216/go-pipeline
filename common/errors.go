package common

import (
	"errors"
)

var (
	ErrOutOfRange = errors.New("out of range")
	ErrIgnore        = errors.New("ignore")
	ErrInvalidFormat = errors.New("invalid format")
	ErrInvalidPara   = errors.New("invalid parameter")
)
