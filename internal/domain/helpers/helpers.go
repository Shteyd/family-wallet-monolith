package helpers

import "monolith/internal/domain"

type Entity interface {
	domain.Customer
}

func IsEmpty[T Entity](entity T) bool {
	var zero T
	return entity == zero
}
