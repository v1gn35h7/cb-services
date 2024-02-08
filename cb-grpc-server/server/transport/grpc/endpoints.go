package grpc

import (
	"github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	"github.com/v1gn35h7/cb-grpc-server/server/service"
)

type grpcEndpoints struct {
	getHealth            endpoint.Endpoint
	bookTicket           endpoint.Endpoint
	getSeatArrangenments endpoint.Endpoint
	removeUser           endpoint.Endpoint
	modifySeat           endpoint.Endpoint
	getReceipt           endpoint.Endpoint
}

func MakeGrpcEndpoints(srvc service.Service, logger log.Logger) grpcEndpoints {

	return grpcEndpoints{
		getHealth:            MakeHealthEndpointMiddleware(srvc, logger),
		bookTicket:           MakeBookTicketEndpointMiddleware(srvc, logger),
		getSeatArrangenments: MakeGetSeatArrangenmentsEndpointMiddleware(srvc, logger),
		removeUser:           MakeRemoveUserEndpointMiddleware(srvc, logger),
		modifySeat:           MakeModifySeatEndpointMiddleware(srvc, logger),
		getReceipt:           MakeGetReceiptEndpointMiddleware(srvc, logger),
	}

}
