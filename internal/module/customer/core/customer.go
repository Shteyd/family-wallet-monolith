package core

import (
	"context"
)

type (
	CustomerUsecase interface {
		Save(context.Context, Customer) (Customer, error)
		GetByCreds(context.Context, Customer) (Customer, error)
	}

	CustomerRepository interface {
		Create(context.Context, Customer) (Customer, error)
		GetById(context.Context, Customer) (Customer, error)
		GetByCreds(context.Context, Customer) (Customer, error)
		Update(context.Context, Customer) error
		UpdateEmailConfirmation(context.Context, Customer) error
		Delete(context.Context, Customer) error
	}
)
