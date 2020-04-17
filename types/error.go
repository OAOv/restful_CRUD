package types

import (
	"errors"
)

var (
	ErrServerQueryError = errors.New("server cannot query")
	ErrInvalidType      = errors.New("invalid type")
	ErrInvalidParams    = errors.New("invalid parameters")
)
