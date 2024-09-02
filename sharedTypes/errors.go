package sharedtypes

import (
	"errors"
)

var (
	ErrorNotFound = errors.New("not found")
	ErrorInternal = errors.New("internal error")
	ErrorInvalid  = errors.New("invalid")
)
