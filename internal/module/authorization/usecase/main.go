package usecase

import (
	"context"
	"monolith/internal/domain"
	"monolith/internal/module/authorization/core"
	"monolith/internal/module/authorization/usecase/internal/utils"
	customer "monolith/internal/module/customer/core"
	password "monolith/internal/module/password/core"
	token "monolith/internal/module/token/core"
	"time"

	"github.com/pkg/errors"
)

type _AuthorizationUsecase struct {
	Logger             domain.Logger
	CustomerRepository customer.CustomerRepository
	PasswordRepository password.PasswordRepository
	TokenRepository    token.TokenRepository

	defaultTimeout time.Duration
}

func NewAuthorizationUsecase(
	logger domain.Logger,
	customerRepository customer.CustomerRepository,
	passwordRepository password.PasswordRepository,
	tokenRepository token.TokenRepository,
	timeout time.Duration,
) core.AuthorizationUsecase {
	return &_AuthorizationUsecase{
		Logger:             logger,
		CustomerRepository: customerRepository,
		PasswordRepository: passwordRepository,
		TokenRepository:    tokenRepository,
		defaultTimeout:     timeout,
	}
}

func (usecase *_AuthorizationUsecase) GenerateCustomerTokens(ctx context.Context, entity customer.Customer) (token.Token, error) {
	accessToken, err := usecase.TokenRepository.GenerateAccessToken(ctx, entity.Id)
	if err != nil {
		return token.Token{}, errors.Wrap(err, "generate access token error")
	}

	refreshToken, err := usecase.TokenRepository.GenerateRefreshToken(ctx, entity.Id)
	if err != nil {
		return token.Token{}, errors.Wrap(err, "generate refresh token error")
	}

	return utils.NewToken(accessToken, refreshToken), nil
}

func (usecase *_AuthorizationUsecase) GenerateAndSetPassword(ctx context.Context, entity *customer.Customer) error {
	password, err := usecase.PasswordRepository.GeneratePassword(ctx, entity.Password)
	if err != nil {
		return errors.Wrap(err, "generate password error")
	}
	entity.Password = password

	return nil
}
