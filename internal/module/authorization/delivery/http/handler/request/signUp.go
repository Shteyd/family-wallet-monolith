package request

import "monolith/internal/module/customer/core"

type SignUpBody struct {
	Username string `json:"username" validate:"omitempty,max=255"`
	Email    string `json:"email" validate:"required,email,max=360"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

func (model SignUpBody) ToEntity() core.Customer {
	return core.Customer{
		Username: model.Username,
		Email:    model.Email,
		Password: model.Password,
	}
}
