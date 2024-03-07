package fibonacci

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUnimplemented       = status.Error(codes.Internal, "unimplemented")
	ErrNotfound            = status.Error(codes.NotFound, "not found")
	ErrTypeAssertionFailed = status.Error(codes.Internal, "type assertion failed")
	ErrTooBig              = status.Error(codes.OutOfRange, "too big")
	ErrInvalidIndex        = status.Error(codes.InvalidArgument, "invalid index, index must be grater than or equal to zero.")
	ErrInvalidNumber       = status.Error(codes.InvalidArgument, "invalid number")
	ErrInvalidBitSize      = status.Error(codes.InvalidArgument, "invalid bit size")
)
