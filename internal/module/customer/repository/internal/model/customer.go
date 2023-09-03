package model

import (
	"database/sql"
	"monolith/internal/module/customer/core"
	"time"
)

type Customer struct {
	Id                int            `db:"id"`
	Username          sql.NullString `db:"username"`
	Email             string         `db:"email"`
	EmailConfirmation bool           `db:"email_confirmation"`
	Password          string         `db:"password"`
	CreatedAt         time.Time      `db:"created_at"`
	UpdatedAt         time.Time      `db:"updated_at"`
	DeletedAt         time.Time      `db:"deleted_at"`
}

func NewCustomer(entity core.Customer) Customer {
	return Customer{
		Id: entity.Id,
		Username: sql.NullString{
			String: entity.Username,
			Valid:  entity.Username != "",
		},
		Email:     entity.Email,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
}

func (model Customer) TableName() string {
	return "customer"
}

func (model Customer) ToEntity() core.Customer {
	return core.Customer{
		Id:        model.Id,
		Username:  model.Username.String,
		Email:     model.Email,
		Password:  model.Password,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
	}
}
