package usecase

import (
	"context"
	"monolith/internal/domain"
	"monolith/internal/module/customer/core"
)

func (usecase *_CustomerUsecase) GetByCreds(ctx context.Context, entity core.Customer) (core.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.defaultContextTimeout)
	defer cancel()

	entity, err := usecase.CustomerRepository.GetByCreds(ctx, entity)
	if err != nil {
		usecase.Logger.Error(err, domain.LoggerArgs{
			"customer_id":    entity.Id,
			"customer_email": entity.Email,
		})
		return core.Customer{}, domain.ErrorInternalServer
	}

	if entity.IsEmpty() {
		usecase.Logger.Info("empty customer", domain.LoggerArgs{
			"customer_id": entity.Id,
		})
		return core.Customer{}, domain.ErrorNotFound
	}

	return entity, nil
}
