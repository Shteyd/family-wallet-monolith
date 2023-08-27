package repository

import (
	"context"
	"monolith/internal/module/customer/adapter/redis"
	"monolith/internal/module/customer/core"
	"monolith/internal/module/customer/repository/internal/model"
	"monolith/internal/module/customer/repository/internal/query"

	"github.com/pkg/errors"
	cache "github.com/redis/go-redis/v9"
)

func (repository *_CustomerRepository) Get(ctx context.Context, entity core.Customer) (core.Customer, error) {
	model := model.NewCustomer(entity)

	if err := repository.RedisAdapter.Get(ctx, redis.GetCustomerKey(model.Id)).Scan(&model); err != nil {
		if !errors.Is(err, cache.Nil) {
			return core.Customer{}, errors.Wrap(err, "get customer from redis error")
		}
	} else {
		return model.ToEntity(), nil
	}

	sql, args, err := query.GetSelect(model)
	if err != nil {
		return core.Customer{}, errors.Wrap(err, "generate select customer sql-query error")
	}

	if err := repository.PostgresAdapter.GetConnect().QueryRow(ctx, sql, args...).Scan(&model); err != nil {
		return core.Customer{}, errors.Wrap(err, "get customer from database error")
	}

	if err := repository.RedisAdapter.Set(ctx, redis.GetCustomerKey(model.Id), model).Err(); err != nil {
		return core.Customer{}, errors.Wrap(err, "set customer in redis error")
	}

	return model.ToEntity(), nil
}
