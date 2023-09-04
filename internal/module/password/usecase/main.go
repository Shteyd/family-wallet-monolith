package usecase

import (
	"context"
	"monolith/internal/domain"
	customer "monolith/internal/module/customer/core"
	"monolith/internal/module/password/core"
	"monolith/internal/module/password/usecase/internal/utils"
	"time"

	"github.com/pkg/errors"
)

type _PasswordUsecase struct {
	Logger             domain.Logger
	CustomerRepository customer.CustomerRepository
	PasswordRepository core.PasswordRepository
	defaultTimeout     time.Duration
}

func NewPasswordUsecase(
	logger domain.Logger,
	customerRepository customer.CustomerRepository,
	passwordRepository core.PasswordRepository,
	timeout time.Duration,
) core.PasswordUsecase {
	return &_PasswordUsecase{
		Logger:             logger,
		CustomerRepository: customerRepository,
		PasswordRepository: passwordRepository,
		defaultTimeout:     timeout,
	}
}

func (usecase *_PasswordUsecase) ChangePassword(ctx context.Context, entity core.Password) error {
	ctx, cancel := context.WithTimeout(ctx, usecase.defaultTimeout)
	defer cancel()

	password, err := usecase.PasswordRepository.GeneratePassword(ctx, entity.Password)
	if err != nil {
		usecase.Logger.Error(errors.Wrap(err, "generate password error"), domain.LoggerArgs{
			"customer_id": entity.CustomerId,
		})
		return domain.ErrorInternalServer
	}
	entity.Password = password

	customer := utils.ConvertPasswordToCustomer(entity)
	if err := usecase.CustomerRepository.UpdatePassword(ctx, customer); err != nil {
		usecase.Logger.Error(errors.Wrap(err, "generate password error"), domain.LoggerArgs{
			"customer_id": entity.CustomerId,
		})
		return domain.ErrorInternalServer
	}

	return nil
}
