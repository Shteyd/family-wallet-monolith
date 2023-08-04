package errors

import (
	"monolith/internal/common/byteconv"

	"github.com/pkg/errors"
)

var (
	As   = errors.As
	Is   = errors.Is
	Wrap = errors.Wrap
)

type Error string

const (
	DatabaseError Error = "database error"
)

func New(err Error) error {
	return errors.New(byteconv.Convert[Error, string](err))
}
