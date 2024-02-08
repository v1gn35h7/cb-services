package service

import (
	"github.com/go-kit/log"
)

/*
* Booking service
* Encapsulates core bussines logic
 */
type Service interface {
	bookingService
}

type service struct {
	logger log.Logger
}

/*
* Create core booking service instance
 */
func New(l log.Logger) service {
	return service{logger: l}
}
