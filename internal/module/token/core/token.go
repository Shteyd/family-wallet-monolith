package core

import "context"

type (
	TokenUsecase interface {
		RefreshToken(context.Context, Token) (Token, error)
	}

	TokenRepository interface {
		GenerateToken(context.Context) (Token, error)
		ValidateToken(context.Context) error
		RefreshToken(context.Context) (Token, error)
	}
)
