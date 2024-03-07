package controllers

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

var (
	errCodeMapping = map[codes.Code]int{
		codes.InvalidArgument:    http.StatusBadRequest,
		codes.NotFound:           http.StatusNotFound,
		codes.OutOfRange:         http.StatusBadRequest,
		codes.FailedPrecondition: http.StatusBadRequest,
	}
)
