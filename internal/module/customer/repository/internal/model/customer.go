package model

import (
	"database/sql"
	"monolith/internal/module/customer/core"
)

type Customer struct {
	Id                int            `db:"id"`
	Username          sql.NullString `db:"username"`
	Email             string         `db:"email"`
	EmailConfirmation bool           `db:"email_confirmation"`
	Password          string         `db:"password_hash"`
	CreatedAt         sql.NullTime   `db:"created_at"`
	UpdatedAt         sql.NullTime   `db:"updated_at"`
	DeletedAt         sql.NullTime   `db:"deleted_at"`
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
		CreatedAt: sql.NullTime{Time: entity.CreatedAt, Valid: true},
		UpdatedAt: sql.NullTime{Time: entity.UpdatedAt, Valid: true},
		DeletedAt: sql.NullTime{Time: entity.DeletedAt, Valid: true},
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
		CreatedAt: model.CreatedAt.Time,
		UpdatedAt: model.UpdatedAt.Time,
		DeletedAt: model.DeletedAt.Time,
	}
}
