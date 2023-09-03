package repository

import (
	"context"
	"monolith/internal/domain"
	"monolith/internal/module/token/adapter/token"
	"monolith/internal/module/token/core"

	"github.com/pkg/errors"
)

type _TokenRepository struct {
	TokenAdapter token.TokenAdapter
}

func NewTokenRepository(tokenAdapter token.TokenAdapter) core.TokenRepository {
	return &_TokenRepository{TokenAdapter: tokenAdapter}
}

func (repository *_TokenRepository) DecodeToken(_ context.Context, token string) (domain.TokenClaims, error) {
	claims, err := repository.TokenAdapter.Parse(token)
	if err != nil {
		return domain.TokenClaims{}, errors.Wrap(err, "parse token error")
	}

	return claims, nil
}

func (repository *_TokenRepository) GenerateAccessToken(_ context.Context, customerId int) (string, error) {
	accessToken, err := repository.TokenAdapter.GenerateAccess(customerId)
	if err != nil {
		return "", errors.Wrap(err, "generate access token error")
	}

	return accessToken, nil
}

func (repository *_TokenRepository) GenerateRefreshToken(_ context.Context, customerId int) (string, error) {
	regreshToken, err := repository.TokenAdapter.GenerateRefresh(customerId)
	if err != nil {
		return "", errors.Wrap(err, "generate refresh token error")
	}

	return regreshToken, nil
}
