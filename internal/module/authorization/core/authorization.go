package core

import (
	"context"
	customer "monolith/internal/module/customer/core"
)

type (
	AuthorizationUsecase interface {
		SignIn(context.Context, customer.Customer) error
		SignUp(context.Context, customer.Customer) error
	}
)
