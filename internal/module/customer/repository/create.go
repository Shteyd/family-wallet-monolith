package repository

import (
	"context"
	"monolith/internal/module/customer/core"
	"monolith/internal/module/customer/repository/internal/model"
	"monolith/internal/module/customer/repository/internal/query"

	"github.com/pkg/errors"
)

func (repository *_CustomerRepository) Create(ctx context.Context, entity core.Customer) (core.Customer, error) {
	insertModel := model.NewCustomer(entity)

	sql, args, err := query.GetInsert(insertModel)
	if err != nil {
		return core.Customer{}, errors.Wrap(err, "generate create customer sql-query error")
	}

	transaction, err := repository.PostgresAdapter.Begin(ctx)
	if err != nil {
		return core.Customer{}, errors.Wrap(err, "begin create customer transaction error")
	}
	connection := transaction.GetConnect()

	model, err := connection.Query(ctx, sql, args...)
	if err != nil {
		if err := transaction.Rollback(ctx); err != nil {
			return core.Customer{}, errors.Wrap(err, "rollback create customer transaction after execute query error")
		}
		return core.Customer{}, errors.Wrap(err, "create customer in database error")
	}

	if err := transaction.Commit(ctx); err != nil {
		return core.Customer{}, errors.Wrap(err, "commit create customer transaction error")
	}

	entity = model.ToEntity()

	repository.CacheAdapter.Set(entity)

	return entity, nil
}
