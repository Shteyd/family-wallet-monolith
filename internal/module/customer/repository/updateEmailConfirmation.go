package repository

import (
	"context"
	"monolith/internal/module/customer/core"
	"monolith/internal/module/customer/repository/internal/model"
	"monolith/internal/module/customer/repository/internal/query"

	"github.com/pkg/errors"
)

func (repository *_CustomerRepository) UpdateEmailConfirmation(ctx context.Context, entity core.Customer) error {
	model := model.NewCustomer(entity)
	sql, args, err := query.GetUpdateEmailConfirmation(model)
	if err != nil {
		return errors.Wrap(err, "generate update email confirmation error")
	}

	connection := repository.PostgresAdapter.GetConnect()
	if _, err := connection.Exec(ctx, sql, args...); err != nil {
		return errors.Wrap(err, "execute update query error")
	}

	return nil
}
