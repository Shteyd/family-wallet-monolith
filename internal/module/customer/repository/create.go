package repository

import (
	"context"
	"monolith/internal/module/customer/adapter/redis"
	"monolith/internal/module/customer/core"
	"monolith/internal/module/customer/repository/internal/model"
	"monolith/internal/module/customer/repository/internal/query"

	"github.com/pkg/errors"
)

func (repository *_CustomerRepository) Create(ctx context.Context, entity core.Customer) (core.Customer, error) {
	model := model.NewCustomer(entity)

	sql, args, err := query.GetInsert(model)
	if err != nil {
		return core.Customer{}, errors.Wrap(err, "generate create customer sql-query error")
	}

	if err := repository.PostgresAdapter.GetConnect().QueryRow(ctx, sql, args...).Scan(&model); err != nil {
		return core.Customer{}, errors.Wrap(err, "create customer in database error")
	}

	if err := repository.RedisAdapter.Set(ctx, redis.GetCustomerKey(model.Id), model).Err(); err != nil {
		// TODO: тут надо бы не error-layer кидать, а warn. Реализовать выбор слоя в apperrors
		return model.ToEntity(), errors.Wrap(err, "set into redis error")
	}

	return model.ToEntity(), nil
}
