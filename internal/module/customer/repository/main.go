package repository

import (
	"context"
	"monolith/internal/domain"
	"monolith/internal/module/customer/adapter/postgres"
	"monolith/internal/module/customer/core"
	"monolith/internal/module/customer/repository/internal/model"
	"monolith/internal/module/customer/repository/internal/query"

	"github.com/pkg/errors"
)

type _CustomerRepository struct {
	DatabaseManager domain.DatabaseManager
}

func NewCustomerRepository(
	databaseManager domain.DatabaseManager,
) core.CustomerRepository {
	return &_CustomerRepository{
		DatabaseManager: databaseManager,
	}
}

func (repository *_CustomerRepository) Create(ctx context.Context, entity core.Customer) (core.Customer, error) {
	model := model.NewCustomer(entity)
	sql, args, err := query.GetCreate(model)
	if err != nil {
		return core.Customer{}, errors.Wrap(err, "generate sql-query error")
	}

	postgresAdapter := postgres.NewPostgresAdapter(repository.DatabaseManager.GetConnect())
	if err := postgresAdapter.QueryRow(ctx, sql, args...).Scan(&model); err != nil {
		return core.Customer{}, errors.Wrap(err, "create customer in database error")
	}

	// TODO: Redis adapter
	// redisAdapter := redis.NewRedisAdapter()
	// if err := redisAdapter.Send(ctx, model); err != nil {
	// 		return core.Customer{}, errors.Wrap(err, "hui")
	// }

	return model.ToEntity(), nil
}

func (repository *_CustomerRepository) Delete(ctx context.Context, entity core.Customer) error {
	panic("unimplemented")
}

func (repository *_CustomerRepository) Get(ctx context.Context, entity core.Customer) (core.Customer, error) {
	panic("unimplemented")
}

func (repository *_CustomerRepository) Update(ctx context.Context, entity core.Customer) error {
	panic("unimplemented")
}
