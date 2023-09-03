package usecase

import (
	"context"
	"monolith/internal/domain"
	customer "monolith/internal/module/customer/core"
	token "monolith/internal/module/token/core"

	"github.com/pkg/errors"
)

func (usecase *_AuthorizationUsecase) SignUp(ctx context.Context, entity customer.Customer) (token.Token, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.defaultTimeout)
	defer cancel()

	if err := usecase.GenerateAndSetPassword(ctx, &entity); err != nil {
		usecase.Logger.Error(errors.Wrap(err, "generate password error"), domain.LoggerArgs{
			"customer_id": entity.Id,
		})
		return token.Token{}, domain.ErrorInternalServer
	}

	customer, err := usecase.CustomerRepository.Create(ctx, entity)
	if err != nil {
		usecase.Logger.Error(errors.Wrap(err, "create customer error"), nil)
		return token.Token{}, domain.ErrorInternalServer
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
