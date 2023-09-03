package repository

import (
	"context"
	"monolith/internal/module/customer/core"
	"monolith/internal/module/customer/repository/internal/model"
	"monolith/internal/module/customer/repository/internal/query"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (repository *_CustomerRepository) GetByCreds(ctx context.Context, entity core.Customer) (core.Customer, error) {
	connection := repository.PostgresAdapter.GetConnect()

	selectModel := model.NewCustomer(entity)

	sql, args, err := query.GetSelectByCreds(selectModel)
	if err != nil {
		return core.Customer{}, errors.Wrap(err, "generate select customer by creds sql query error")
	}

	rows, err := connection.Query(ctx, sql, args...)
	if err != nil {
		return core.Customer{}, errors.Wrap(err, "execute select customer by creds sql query error")
	}

	model, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByNameLax[model.Customer])
	if err != nil {
		return core.Customer{}, errors.Wrap(err, "scan customer model error")
	}

	return model.ToEntity(), nil
}
