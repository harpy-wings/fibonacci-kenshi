package fibonacci

import "errors"

var (
	ErrUnimplemented       = errors.New("unimplemented")
	ErrNotfound            = errors.New("not found")
	ErrTypeAssertionFailed = errors.New("type assertion failed")
	ErrTooBig              = errors.New("too big")
	ErrInvalidIndex        = errors.New("invalid index, index must be grater than or equal to zero.")
	ErrInvalidNumber       = errors.New("invalid number")
	ErrInvalidBitSize      = errors.New("invalid bit size")
)
