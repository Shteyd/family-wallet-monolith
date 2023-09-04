package repository

import (
	"context"
	"monolith/internal/module/password/adapter/cryptohash"
	"monolith/internal/module/password/core"

	"github.com/pkg/errors"
)

type _PasswordRepository struct {
	CryptoHashAdapter cryptohash.CryptoHashAdapter
}

func NewPasswordRepository(
	cryptoHashAdapter cryptohash.CryptoHashAdapter,
) core.PasswordRepository {
	return &_PasswordRepository{
		CryptoHashAdapter: cryptoHashAdapter,
	}
}

func (repository *_PasswordRepository) GeneratePassword(ctx context.Context, password string) (string, error) {
	password, err := repository.CryptoHashAdapter.GeneratePasswordHash(password)
	if err != nil {
		return "", errors.Wrap(err, "generate password hash error")
	}

	return password, nil
}
