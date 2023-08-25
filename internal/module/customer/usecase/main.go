package usecase

import (
	"monolith/internal/domain"
	"monolith/internal/module/customer/core"
	"time"
)

type _CustomerUsecase struct {
	Logger             domain.Logger
	CustomerRepository core.CustomerRepository

	defaultContextTimeout time.Duration
}

func NewCustomerUsecase(
	logger domain.Logger,
	customerRepository core.CustomerRepository,
	contextTimeout time.Duration,
) core.CustomerUsecase {
	return &_CustomerUsecase{
		Logger:                logger,
		CustomerRepository:    customerRepository,
		defaultContextTimeout: contextTimeout,
	}
}
