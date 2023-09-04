package request

import "monolith/internal/module/customer/core"

type SignInBody struct {
	Email    string `json:"email" validate:"required,email,max=360"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

func (model SignInBody) ToEntity() core.Customer {
	return core.Customer{
		Email:    model.Email,
		Password: model.Password,
	}
}
