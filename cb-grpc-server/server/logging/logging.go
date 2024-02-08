package logging

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/log"
)

// Endpoint middleware
func LoggingMiddleware(logger kitlog.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling gRPC endpoint")
			defer logger.Log("msg", "called gRPC endpoint")
			return next(ctx, request)
		}
	}
}
