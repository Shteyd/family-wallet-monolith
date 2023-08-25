package usecase

import (
	"context"
	"monolith/internal/domain"
)

func (usecase *_CustomerUsecase) Update(ctx context.Context, entity domain.Customer) error {
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
