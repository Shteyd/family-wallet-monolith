package usecase

import (
	"monolith/internal/domain"
	"time"
)

type _CustomerUsecase struct {
	Logger             domain.Logger
	CustomerRepository domain.CustomerRepository

	defaultContextTimeout time.Duration
}

func NewCustomerUsecase(
	logger domain.Logger,
	customerRepository domain.CustomerRepository,
	contextTimeout time.Duration,
) domain.CustomerUsecase {
	return &_CustomerUsecase{
		Logger:                logger,
		CustomerRepository:    customerRepository,
		defaultContextTimeout: contextTimeout,
	}
}
