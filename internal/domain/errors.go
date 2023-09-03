package domain

import "errors"

var (
	ErrorInternalServer = errors.New("interal server error")
	ErrorNotFound       = errors.New("not found error")
	ErrorInvalidType    = errors.New("invalid type error")
)
