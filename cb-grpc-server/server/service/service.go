package service

import (
	"github.com/go-kit/log"
)

type Service interface {
	bookingService
}

type service struct {
	logger log.Logger
}

func New(l log.Logger) service {
	return service{logger: l}
}
