package repository

import (
	"context"
	"monolith/internal/domain"
	"monolith/internal/module/customer/core"
	"monolith/internal/module/customer/repository/internal/model"
	"monolith/internal/module/customer/repository/internal/query"

	"github.com/pkg/errors"
)

func (repository *_CustomerRepository) GetById(ctx context.Context, entity core.Customer) (core.Customer, error) {
	if entity, err := repository.CacheAdapter.Get(entity); err != nil {
		if !errors.Is(err, domain.ErrorNotFound) {
			return core.Customer{}, errors.Wrap(err, "cache customer error")
		}
	} else {
		return entity, nil
	}

	selectModel := model.NewCustomer(entity)
	sql, args, err := query.GetSelectById(selectModel)
	if err != nil {
		return core.Customer{}, errors.Wrap(err, "generate select customer sql query error")
	}

	connection := repository.PostgresAdapter.GetConnect()

	model, err := connection.Query(ctx, sql, args...)
	if err != nil {
		return core.Customer{}, errors.Wrap(err, "execute select customer query error")
	}

	return model.ToEntity(), nil
}
