package types

import (
	"errors"
)

var (
	ErrServerQueryError  = errors.New("server cannot query")
	ErrInvalidType       = errors.New("invalid type")
	ErrInvalidParams     = errors.New("invalid parameters")
	ErrInvalidInputRange = errors.New("invaild input range")
	ErrNotFound          = errors.New("paramter is not in list")
	ErrEmptyInput        = errors.New("empty input")
)
