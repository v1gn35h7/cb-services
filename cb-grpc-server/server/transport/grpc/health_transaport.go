package grpc

import (
	"context"

	"github.com/v1gn35h7/cb-grpc-server/server/pb"
)

// Health service transport
func (s *grpcServer) GetHealth(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) {
	_, rep, err := s.getHealth.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	res, ok := rep.(*pb.HealthResponse)

	if ok != true {
		res = &pb.HealthResponse{Status: "down"}
	}
	return res, nil
}
