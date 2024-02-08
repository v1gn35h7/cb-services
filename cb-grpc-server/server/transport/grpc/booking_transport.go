package grpc

import (
	"context"

	"github.com/v1gn35h7/cb-grpc-server/server/pb"
)

// Booking serviec transport
func (s *grpcServer) BookTicket(ctx context.Context, req *pb.BookingRequest) (*pb.BookingResponse, error) {
	_, rep, err := s.bookTicket.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	res, ok := rep.(*pb.BookingResponse)

	if ok != true {
		res = &pb.BookingResponse{}
	}
	return res, nil
}

func (s *grpcServer) RemoveUser(ctx context.Context, req *pb.RemoveUserRequest) (*pb.RemoveUserResponse, error) {
	_, rep, err := s.removeUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	res, ok := rep.(*pb.RemoveUserResponse)

	if ok != true {
		res = &pb.RemoveUserResponse{}
	}
	return res, nil
}

func (s *grpcServer) ModifySeat(ctx context.Context, req *pb.ModifySeatRequest) (*pb.ModifySeatResponse, error) {
	_, rep, err := s.modifySeat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	res, ok := rep.(*pb.ModifySeatResponse)

	if ok != true {
		res = &pb.ModifySeatResponse{}
	}
	return res, nil
}

func (s *grpcServer) GetReceipt(ctx context.Context, req *pb.ReceiptRequest) (*pb.BookingResponse, error) {
	_, rep, err := s.getReceipt.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	res, ok := rep.(*pb.BookingResponse)

	if ok != true {
		res = &pb.BookingResponse{}
	}
	return res, nil
}

func (s *grpcServer) GetSeatArrangenments(ctx context.Context, req *pb.SeatArrangmentRequest) (*pb.SeatArrangmentResponse, error) {
	_, rep, err := s.getSeatArrangenments.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SeatArrangmentResponse), nil
}
