package usecase

import (
	"context"
	"monolith/internal/domain"
	"monolith/internal/module/customer/core"
)

func (usecase *_CustomerUsecase) Create(ctx context.Context, entity core.Customer) (core.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.defaultContextTimeout)
	defer cancel()

	entity, err := usecase.CustomerRepository.Create(ctx, entity)
	if err != nil {
		usecase.Logger.Error(err, nil)
		return core.Customer{}, domain.ErrorInternalServer
	}

	return entity, nil
}
