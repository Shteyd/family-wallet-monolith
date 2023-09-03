package core

import (
	"context"
	"monolith/internal/domain"
)

type (
	TokenUsecase interface {
		RefreshToken(context.Context, Token) (Token, error)
	}

	TokenRepository interface {
		GenerateAccessToken(context.Context, int) (string, error)
		GenerateRefreshToken(context.Context, int) (string, error)
		DecodeToken(context.Context, string) (domain.TokenClaims, error)
	}
)
