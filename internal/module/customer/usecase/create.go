package usecase

import (
	"context"
	"monolith/internal/domain"
)

func (usecase *_CustomerUsecase) Create(ctx context.Context, entity domain.Customer) (domain.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.defaultContextTimeout)
	defer cancel()

	entity, err := usecase.CustomerRepository.Create(ctx, entity)
	if err != nil {
		usecase.Logger.Error(err, nil)
		return domain.Customer{}, domain.ErrorInternalServer
	}

	return entity, nil
}
