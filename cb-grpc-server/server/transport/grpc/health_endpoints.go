package grpc

import (
	"context"

	log "github.com/go-kit/kit/log"

	"github.com/go-kit/kit/endpoint"
	"github.com/v1gn35h7/cb-grpc-server/server/logging"
	"github.com/v1gn35h7/cb-grpc-server/server/pb"
	"github.com/v1gn35h7/cb-grpc-server/server/service"
)

// Endpoints creators
func makeHealthEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		h := &pb.HealthResponse{}
		h.Status = "OK"

		return h, nil
	}
}

func MakeHealthEndpointMiddleware(srvc service.Service, logger log.Logger) endpoint.Endpoint {
	healthEndpoint := makeHealthEndpoint(srvc)
	return logging.LoggingMiddleware(logger)(healthEndpoint)
}
