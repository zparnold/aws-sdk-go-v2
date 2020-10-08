package v4

import (
	"context"

	"github.com/awslabs/smithy-go/middleware"
)

func NewPresignHTTPRequestMiddleware() *PresignHTTPRequestMiddleware {
	return &PresignHTTPRequestMiddleware{}
}

type PresignHTTPRequestMiddleware struct{}

func (*PresignHTTPRequestMiddleware) ID() string { return "PresignHTTPRequestMiddleware" }

// HandleFinalize will take the provided input and sign the request using the
// SigV4 presign authentication scheme.
//
// Since the signed request is not a valid HTTP request
func (s *PresignHTTPRequestMiddleware) HandleFinalize(
	ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler,
) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {

	return next.HandleFinalize(ctx, in)
}
