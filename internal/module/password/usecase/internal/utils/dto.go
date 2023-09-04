package utils

import (
	customer "monolith/internal/module/customer/core"
	password "monolith/internal/module/password/core"
)

func ConvertPasswordToCustomer(entity password.Password) customer.Customer {
	return customer.Customer{
		Id:       entity.CustomerId,
		Password: entity.Password,
	}
}
