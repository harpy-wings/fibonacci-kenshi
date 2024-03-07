package codec

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidCodec = status.Error(codes.InvalidArgument, "invalid codec")
)
