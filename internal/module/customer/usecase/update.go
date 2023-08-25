package usecase

import (
	"context"
	"monolith/internal/domain"
	"monolith/internal/module/customer/core"
)

func (usecase *_CustomerUsecase) Update(ctx context.Context, entity core.Customer) error {
	ctx, cancel := context.WithTimeout(ctx, usecase.defaultContextTimeout)
	defer cancel()

	if err := usecase.CustomerRepository.Update(ctx, entity); err != nil {
		usecase.Logger.Error(err, domain.LoggerArgs{
			"customer_id":    entity.Id,
			"customer_email": entity.Email,
		})
		return domain.ErrorInternalServer
	}

	return nil
}
