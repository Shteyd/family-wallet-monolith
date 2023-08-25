package core

import (
	"context"
)

type (
	CustomerUsecase interface {
		Create(context.Context, Customer) (Customer, error)
		Get(context.Context, Customer) (Customer, error)
		Update(context.Context, Customer) error
	}

	CustomerRepository interface {
		Create(context.Context, Customer) (Customer, error)
		Get(context.Context, Customer) (Customer, error)
		Update(context.Context, Customer) error
		Delete(context.Context, Customer) error
	}
)
