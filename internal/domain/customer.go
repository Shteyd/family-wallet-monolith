package domain

import (
	"context"
	"time"
)

type Customer struct {
	Id        int       `json:"id" example:"1"`
	Username  string    `json:"username" example:"Shteyd"`
	Email     string    `json:"email" example:"example@mail.com"`
	Password  string    `json:"password" example:"qwerty123"`
	CreatedAt time.Time `json:"created_at" example:"2023-08-25 17:27:35.811169+00"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-08-25 17:27:35.811169+00"`
	DeletedAt time.Time `json:"deleted_at" example:"2023-08-25 17:27:35.811169+00"`
}

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
