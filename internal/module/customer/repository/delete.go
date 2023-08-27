package repository

import (
	"context"
	"monolith/internal/module/customer/adapter/redis"
	"monolith/internal/module/customer/core"
	"monolith/internal/module/customer/repository/internal/model"
	"monolith/internal/module/customer/repository/internal/query"

	"github.com/pkg/errors"
)

func (repository *_CustomerRepository) Delete(ctx context.Context, entity core.Customer) error {
	model := model.NewCustomer(entity)

	if err := repository.RedisAdapter.Del(ctx, redis.GetCustomerKey(model.Id)).Err(); err != nil {
		return errors.Wrap(err, "delete from redis error")
	}

	sql, args, err := query.GetDelete(model)
	if err != nil {
		return errors.Wrap(err, "generate delete customer sql-query error")
	}

	if _, err := repository.PostgresAdapter.GetConnect().Exec(ctx, sql, args...); err != nil {
		return errors.Wrap(err, "delete customer from database error")
	}

	return nil
}
