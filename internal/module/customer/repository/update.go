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
	if err := repository.RedisAdapter.Del(ctx, redis.GetCustomerKey(entity.Id)).Err(); err != nil {
		return errors.Wrap(err, "remove customer from redis error")
	}

	model := model.NewCustomer(entity)
	sql, args, err := query.GetUpdate(model)
	if err != nil {
		return errors.Wrap(err, "generate update customer sql-query error")
	}

	connection := repository.PostgresAdapter.GetConnect()
	if _, err := connection.Exec(ctx, sql, args...); err != nil {
		return errors.Wrap(err, "execute update customer query error")
	}

	return nil
}
