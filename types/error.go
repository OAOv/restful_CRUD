package types

import (
	"errors"
)

var (
	ErrServerQueryError = errors.New("server cannot query")
	ErrInvalidType      = errors.New("invalid parameters or struct type")
	ErrInvalidParms     = errors.New("invalid parameters")
)
