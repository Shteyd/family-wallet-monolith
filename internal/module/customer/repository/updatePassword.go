package repository

import (
	"context"
	"monolith/internal/module/customer/core"
	"monolith/internal/module/customer/repository/internal/model"
	"monolith/internal/module/customer/repository/internal/query"

	"github.com/pkg/errors"
)

func (repository *_CustomerRepository) UpdatePassword(ctx context.Context, entity core.Customer) error {
	model := model.NewCustomer(entity)

	sql, args, err := query.GetUpdatePassword(model)
	if err != nil {
		return errors.Wrap(err, "generate update customer password sql query error")
	}

	connection := repository.PostgresAdapter.GetConnect()
	if err := connection.Exec(ctx, sql, args...); err != nil {
		return errors.Wrap(err, "execute update customer query error")
	}

	return nil
}
