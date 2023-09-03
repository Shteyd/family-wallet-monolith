package usecase

import (
	"context"
	"monolith/internal/domain"
	customer "monolith/internal/module/customer/core"
	token "monolith/internal/module/token/core"

	"github.com/pkg/errors"
)

func (usecase *_AuthorizationUsecase) SignIn(ctx context.Context, entity customer.Customer) (token.Token, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.defaultTimeout)
	defer cancel()

	if err := usecase.GenerateAndSetPassword(ctx, &entity); err != nil {
		usecase.Logger.Error(errors.Wrap(err, "generate password error"), domain.LoggerArgs{
			"customer_id": entity.Id,
		})
		return token.Token{}, domain.ErrorInternalServer
	}

	customer, err := usecase.CustomerRepository.GetByCreds(ctx, entity)
	if err != nil {
		usecase.Logger.Error(errors.Wrap(err, "get customer error"), nil)
		return token.Token{}, domain.ErrorInternalServer
	}

	if customer.IsEmpty() {
		usecase.Logger.Debug("customer not found", domain.LoggerArgs{
			"customer_email":    entity.Email,
			"customer_password": entity.Password,
		})
		return token.Token{}, domain.ErrorNotFound
	}

	tokenModel, err := usecase.GenerateCustomerTokens(ctx, customer)
	if err != nil {
		usecase.Logger.Error(errors.Wrap(err, "generate tokens error"), domain.LoggerArgs{
			"customer_id": customer.Id,
		})
		return token.Token{}, domain.ErrorInternalServer
	}

	return tokenModel, nil
}
