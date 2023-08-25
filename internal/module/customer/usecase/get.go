package usecase

import (
	"context"
	"monolith/internal/domain"
	"monolith/internal/domain/helpers"
)

func (usecase *_CustomerUsecase) Get(ctx context.Context, entity domain.Customer) (domain.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.defaultContextTimeout)
	defer cancel()

	entity, err := usecase.CustomerRepository.Get(ctx, entity)
	if err != nil {
		usecase.Logger.Error(err, domain.LoggerArgs{
			"customer_id":    entity.Id,
			"customer_email": entity.Email,
		})
		return domain.Customer{}, domain.ErrorInternalServer
	}

	if helpers.IsEmpty[domain.Customer](entity) {
		usecase.Logger.Warn("empty customer", domain.LoggerArgs{
			"customer_id":    entity.Id,
			"customer_email": entity.Email,
		})
		return domain.Customer{}, domain.ErrorNotFound
	}

	return entity, nil
}
