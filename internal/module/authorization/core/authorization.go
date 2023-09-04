package core

import (
	"context"
	customer "monolith/internal/module/customer/core"
	token "monolith/internal/module/token/core"
)

type (
	AuthorizationUsecase interface {
		SignIn(context.Context, customer.Customer) (token.Token, error)
		SignUp(context.Context, customer.Customer) (token.Token, error)
	}
)
