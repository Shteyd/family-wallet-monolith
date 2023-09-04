package token

import (
	"monolith/internal/domain"
)

type TokenAdapter interface {
	GenerateAccess(customerId int) (string, error)
	GenerateRefresh(customerId int) (string, error)
	Parse(token string) (domain.TokenClaims, error)
}

type _TokenAdapter struct {
	Manager domain.TokenManager
}

func NewTokenAdapter(manager domain.TokenManager) TokenAdapter {
	return &_TokenAdapter{Manager: manager}
}

func (adapter *_TokenAdapter) GenerateAccess(customerId int) (string, error) {
	return adapter.Manager.GenerateAccess(customerId)
}

func (adapter *_TokenAdapter) GenerateRefresh(customerId int) (string, error) {
	return adapter.Manager.GenerateRefresh(customerId)
}

func (adapter *_TokenAdapter) Parse(token string) (domain.TokenClaims, error) {
	return adapter.Manager.Parse(token)
}
