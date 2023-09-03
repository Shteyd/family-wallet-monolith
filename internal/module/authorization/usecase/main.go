package usecase

import (
	"context"
	"monolith/internal/domain"
	"monolith/internal/module/authorization/core"
	customer "monolith/internal/module/customer/core"
	password "monolith/internal/module/password/core"
	"time"

	"github.com/pkg/errors"
)

type _AuthorizationUsecase struct {
	Logger             domain.Logger
	CustomerRepository customer.CustomerRepository
	PasswordRepository password.PasswordRepository

	defaultTimeout time.Duration
}

func NewAuthorizationUsecase(
	logger domain.Logger,
	customerRepository customer.CustomerRepository,
	passwordRepository password.PasswordRepository,
	timeout time.Duration,
) core.AuthorizationUsecase {
	return &_AuthorizationUsecase{
		Logger:             logger,
		CustomerRepository: customerRepository,
		PasswordRepository: passwordRepository,
		defaultTimeout:     timeout,
	}
}

func (usecase *_AuthorizationUsecase) SignIn(ctx context.Context, entity customer.Customer) error {
	ctx, cancel := context.WithTimeout(ctx, usecase.defaultTimeout)
	defer cancel()

	password, err := usecase.PasswordRepository.GeneratePassword(ctx, entity.Password)
	if err != nil {
		usecase.Logger.Error(errors.Wrap(err, "generate password error"), nil)
		return domain.ErrorInternalServer
	}
	entity.Password = password

	customer, err := usecase.CustomerRepository.GetByCreds(ctx, entity)
	if err != nil {
		usecase.Logger.Error(errors.Wrap(err, "get customer error"), nil)
		return domain.ErrorInternalServer
	}

	if customer.IsEmpty() {
		usecase.Logger.Debug("customer not found", domain.LoggerArgs{
			"customer_email":    entity.Email,
			"customer_password": entity.Password,
		})
		return domain.ErrorNotFound
	}

	// TODO: generate token

	return nil
}

func (usecase *_AuthorizationUsecase) SignUp(ctx context.Context, entity customer.Customer) error {
	ctx, cancel := context.WithTimeout(ctx, usecase.defaultTimeout)
	defer cancel()

	password, err := usecase.PasswordRepository.GeneratePassword(ctx, entity.Password)
	if err != nil {
		usecase.Logger.Error(errors.Wrap(err, "generate password error"), nil)
		return domain.ErrorInternalServer
	}
	entity.Password = password

	customer, err := usecase.CustomerRepository.Create(ctx, entity)
	if err != nil {
		usecase.Logger.Error(errors.Wrap(err, "create customer error"), nil)
		return domain.ErrorInternalServer
	}
	_ = customer

	// TODO: generate token

	return nil
}
