package grpc

import (
	"context"
	"strconv"

	log "github.com/go-kit/kit/log"

	"github.com/go-kit/kit/endpoint"
	"github.com/v1gn35h7/cb-grpc-server/server/logging"
	"github.com/v1gn35h7/cb-grpc-server/server/pb"
	"github.com/v1gn35h7/cb-grpc-server/server/service"
)

// Booking service endpoints creators
func makeBookTicketEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(*pb.BookingRequest)

		res, err := srvc.BookTicket(*req)

		if err != nil {
			return nil, err
		}

		pbResponse := pb.BookingResponse{
			User: &pb.User{
				ID:        res.User.ID,
				Email:     res.User.Email,
				FirstName: res.User.FirstName,
				LastName:  res.User.LastName,
			},
			From:  res.From,
			To:    res.To,
			Price: res.Price,
		}

		return pbResponse, err
	}
}

func MakeBookTicketEndpointMiddleware(srvc service.Service, logger log.Logger) endpoint.Endpoint {
	makeBookTicketEndpoint := makeBookTicketEndpoint(srvc)
	return logging.LoggingMiddleware(logger)(makeBookTicketEndpoint)
}

func makeRemoveUserEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(*pb.RemoveUserRequest)

		res, err := srvc.RemoveUser(req)

		if err != nil {
			return nil, err
		}

		pbResponse := &pb.RemoveUserResponse{
			Status: res.Status,
		}

		return pbResponse, err
	}
}

func MakeRemoveUserEndpointMiddleware(srvc service.Service, logger log.Logger) endpoint.Endpoint {
	makeRemoveUserEndpoint := makeRemoveUserEndpoint(srvc)
	return logging.LoggingMiddleware(logger)(makeRemoveUserEndpoint)
}

func makeModifySeatEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(*pb.ModifySeatRequest)

		res, err := srvc.ModifySeat(req)

		if err != nil {
			return nil, err
		}

		pbResponse := &pb.ModifySeatResponse{
			Status: res.Status,
		}

		return pbResponse, err
	}
}

func MakeModifySeatEndpointMiddleware(srvc service.Service, logger log.Logger) endpoint.Endpoint {
	makeSeatEndpoint := makeModifySeatEndpoint(srvc)
	return logging.LoggingMiddleware(logger)(makeSeatEndpoint)
}

func makeGetReceiptEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(*pb.ReceiptRequest)

		res, err := srvc.GetReceipt(req)

		if err != nil {
			return nil, err
		}

		pbResponse := &pb.BookingResponse{
			User: &pb.User{
				ID:        res.User.ID,
				Email:     res.User.Email,
				FirstName: res.User.FirstName,
				LastName:  res.User.LastName,
			},
			From:  res.From,
			To:    res.To,
			Price: res.Price,
		}

		return pbResponse, err
	}
}

func MakeGetReceiptEndpointMiddleware(srvc service.Service, logger log.Logger) endpoint.Endpoint {
	makeReceiptEndpoint := makeGetReceiptEndpoint(srvc)
	return logging.LoggingMiddleware(logger)(makeReceiptEndpoint)
}

func makeGetSeatArrangenmentsEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(*pb.SeatArrangmentRequest)

		res, err := srvc.GetSeatArrangenments(req)

		if err != nil {
			return nil, err
		}

		rpcResponse := &pb.SeatArrangmentResponse{
			Sections: make([]*pb.Section, 0),
		}

		rpcResponse.Sections = append(rpcResponse.Sections, &pb.Section{
			Name:  res.Name,
			Seats: make([]*pb.SeatDetails, 0),
		})

		for _, v := range res.Seats {
			sec := &pb.SeatDetails{
				Number: strconv.Itoa(int(v.ID)),
				UserID: strconv.Itoa(int(v.UserId)),
			}
			rpcResponse.Sections[0].Seats = append(rpcResponse.Sections[0].Seats, sec)
		}

		return rpcResponse, nil
	}
}

func MakeGetSeatArrangenmentsEndpointMiddleware(srvc service.Service, logger log.Logger) endpoint.Endpoint {
	getScriptEndpoint := makeGetSeatArrangenmentsEndpoint(srvc)
	return logging.LoggingMiddleware(logger)(getScriptEndpoint)
}
