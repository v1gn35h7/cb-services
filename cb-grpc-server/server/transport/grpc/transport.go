package grpc

import (
	"context"
	"fmt"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/v1gn35h7/cb-grpc-server/server/pb"
)

// Grpc server
type grpcServer struct {
	getHealth            grpctransport.Handler
	bookTicket           grpctransport.Handler
	getSeatArrangenments grpctransport.Handler
	removeUser           grpctransport.Handler
	modifySeat           grpctransport.Handler
	getReceipt           grpctransport.Handler
	pb.UnimplementedCbServiceServer
}

/*
* Registers all gRPC transports for booking service
 */
func NewGRPCServer(endpoints grpcEndpoints) *grpcServer {

	return &grpcServer{
		getHealth: grpctransport.NewServer(
			endpoints.getHealth,
			// Decodes request
			func(ctx context.Context, i interface{}) (interface{}, error) {
				req, ok := i.(pb.HealthRequest)

				if !ok {
					req = pb.HealthRequest{}
				}

				return req, nil
			},
			// Encodes response
			func(ctx context.Context, i interface{}) (response interface{}, err error) {
				r, ok := i.(*pb.HealthResponse)

				if !ok {
					return nil, fmt.Errorf("Error decoding response")
				}

				return r, nil
			},
		),
		bookTicket: grpctransport.NewServer(
			endpoints.bookTicket,
			func(ctx context.Context, i interface{}) (interface{}, error) {
				request, ok := i.(*pb.BookingRequest)

				if !ok {
					request = &pb.BookingRequest{}
				}

				return request, nil
			},
			func(ctx context.Context, i interface{}) (interface{}, error) {
				r, ok := i.(pb.BookingResponse)

				if !ok {
					return nil, fmt.Errorf("Error decoding response")
				}

				return &r, nil
			},
		),
		getSeatArrangenments: grpctransport.NewServer(
			endpoints.getSeatArrangenments,
			func(ctx context.Context, i interface{}) (interface{}, error) {
				request, ok := i.(*pb.SeatArrangmentRequest)

				if !ok {
					request = &pb.SeatArrangmentRequest{}
				}

				return request, nil
			},
			func(ctx context.Context, i interface{}) (interface{}, error) {
				r, ok := i.(*pb.SeatArrangmentResponse)

				if !ok {
					return nil, fmt.Errorf("Error decoding response")
				}

				return r, nil
			},
		),
		removeUser: grpctransport.NewServer(
			endpoints.removeUser,
			func(ctx context.Context, i interface{}) (interface{}, error) {
				request, ok := i.(*pb.RemoveUserRequest)

				if !ok {
					request = &pb.RemoveUserRequest{}
				}

				return request, nil
			},
			func(ctx context.Context, i interface{}) (interface{}, error) {
				r, ok := i.(*pb.RemoveUserResponse)

				if !ok {
					return nil, fmt.Errorf("Error decoding response")
				}

				return r, nil
			},
		),
		modifySeat: grpctransport.NewServer(
			endpoints.modifySeat,
			func(ctx context.Context, i interface{}) (interface{}, error) {
				request, ok := i.(*pb.ModifySeatRequest)

				if !ok {
					request = &pb.ModifySeatRequest{}
				}

				return request, nil
			},
			func(ctx context.Context, i interface{}) (interface{}, error) {
				r, ok := i.(*pb.ModifySeatResponse)

				if !ok {
					return nil, fmt.Errorf("Error decoding response")
				}

				return r, nil
			},
		),
		getReceipt: grpctransport.NewServer(
			endpoints.getReceipt,
			func(ctx context.Context, i interface{}) (interface{}, error) {
				request, ok := i.(*pb.ReceiptRequest)

				if !ok {
					request = &pb.ReceiptRequest{}
				}

				return request, nil
			},
			func(ctx context.Context, i interface{}) (interface{}, error) {
				r, ok := i.(*pb.BookingResponse)

				if !ok {
					return nil, fmt.Errorf("Error decoding response")
				}

				return r, nil
			},
		),
	}
}
