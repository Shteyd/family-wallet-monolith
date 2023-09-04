package usecase

import (
	"context"
	"monolith/internal/domain"
	"monolith/internal/module/customer/core"

	"github.com/pkg/errors"
)

func (usecase *_CustomerUsecase) Save(ctx context.Context, entity core.Customer) (core.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.defaultContextTimeout)
	defer cancel()

	if entity.Id != 0 {
		if err := usecase.CustomerRepository.Update(ctx, entity); err != nil {
			usecase.Logger.Error(errors.Wrap(err, "update customer error"), domain.LoggerArgs{
				"customer_id": entity.Id,
			})
			return core.Customer{}, domain.ErrorInternalServer
		}

		entity, err := usecase.CustomerRepository.GetById(ctx, entity)
		if err != nil {
			usecase.Logger.Error(errors.Wrap(err, "get customer by id error"), domain.LoggerArgs{
				"customer_id": entity.Id,
			})
			return core.Customer{}, domain.ErrorInternalServer
		}

		return entity, nil
	}

	entity, err := usecase.CustomerRepository.Create(ctx, entity)
	if err != nil {
		usecase.Logger.Error(errors.Wrap(err, "create customer error"), nil)
		return core.Customer{}, domain.ErrorInternalServer
	}

	return entity, nil
}
