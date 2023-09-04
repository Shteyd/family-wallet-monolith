package usecase

import (
	"context"
	"monolith/internal/domain"
	"monolith/internal/module/token/core"
	"monolith/internal/module/token/usecase/internal/utils"
	"time"

	"github.com/pkg/errors"
)

type _TokenUsecase struct {
	Logger          domain.Logger
	TokenRepository core.TokenRepository

	defaultTimeout time.Duration
}

func NewTokenUsecase(
	logger domain.Logger,
	tokenRepository core.TokenRepository,
	defaultTimeout time.Duration,
) core.TokenUsecase {
	return &_TokenUsecase{
		Logger:          logger,
		TokenRepository: tokenRepository,
		defaultTimeout:  defaultTimeout,
	}
}

func (usecase *_TokenUsecase) RefreshToken(ctx context.Context, entity core.Token) (core.Token, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.defaultTimeout)
	defer cancel()

	decodeToken, err := usecase.TokenRepository.DecodeToken(ctx, entity.RefreshToken)
	if err != nil {
		usecase.Logger.Debug(err.Error(), nil)
		return core.Token{}, domain.ErrorInternalServer
	}

	// check expiration time
	if utils.CheckRefreshExpirationTime(decodeToken.Expiration) {
		refreshToken, err := usecase.TokenRepository.GenerateRefreshToken(ctx, decodeToken.CustomerId)
		if err != nil {
			usecase.Logger.Error(errors.Wrap(err, "generate access token error"), domain.LoggerArgs{
				"customer_id": decodeToken.CustomerId,
			})
			return core.Token{}, domain.ErrorInternalServer
		}

		entity.RefreshToken = refreshToken
	}

	accessToken, err := usecase.TokenRepository.GenerateAccessToken(ctx, decodeToken.CustomerId)
	if err != nil {
		usecase.Logger.Error(errors.Wrap(err, "generate access token error"), domain.LoggerArgs{
			"customer_id": decodeToken.CustomerId,
		})
		return core.Token{}, domain.ErrorInternalServer
	}

	entity.AccessToken = accessToken

	return entity, nil
}
