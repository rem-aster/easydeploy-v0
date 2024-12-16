package solution

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	InternalError = status.Error(
		codes.Internal,
		"something went wrong",
	)
)
