package repository

import (
	"context"
	"monolith/internal/module/customer/core"
	"monolith/internal/module/customer/repository/internal/model"
	"monolith/internal/module/customer/repository/internal/query"

	"github.com/pkg/errors"
)

func (repository *_CustomerRepository) Delete(ctx context.Context, entity core.Customer) error {
	connection := repository.PostgresAdapter.GetConnect()

	repository.CacheAdapter.Del(entity)

	deleteModel := model.NewCustomer(entity)
	sql, args, err := query.GetDelete(deleteModel)
	if err != nil {
		return errors.Wrap(err, "generate delete customer sql-query error")
	}

	if err := connection.Exec(ctx, sql, args...); err != nil {
		return errors.Wrap(err, "delete customer from database error")
	}

	return nil
}
