package repository

import (
	"context"
	"monolith/internal/module/customer/adapter/redis"
	"monolith/internal/module/customer/core"
	"monolith/internal/module/customer/repository/internal/model"
	"monolith/internal/module/customer/repository/internal/query"

	"github.com/pkg/errors"
)

func (repository *_CustomerRepository) Update(ctx context.Context, entity core.Customer) error {
	oldEntity, err := repository.Get(ctx, entity)
	if err != nil {
		return errors.Wrap(err, "get customer error")
	}

	oldModel := model.NewCustomer(oldEntity)
	model := model.NewCustomer(entity)

	sql, args, err := query.GetUpdate(oldModel, model)
	if err != nil {
		return errors.Wrap(err, "generate update customer sql-query error")
	}

	transaction, err := repository.PostgresAdapter.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "begin transaction for update customer error")
	}

	if err := transaction.GetConnect().QueryRow(ctx, sql, args...).Scan(&model); err != nil {
		if err := transaction.Rollback(ctx); err != nil {
			return errors.Wrap(err, "rollback customer error")
		}

		return errors.Wrap(err, "update customer id database error")
	}

	if err := repository.RedisAdapter.Del(ctx, redis.GetCustomerKey(model.Id)).Err(); err != nil {
		return errors.Wrap(err, "remove customer from cache error")
	}

	if err := transaction.Commit(ctx); err != nil {
		return errors.Wrap(err, "commit transaction for update customer error")
	}

	return nil
}
