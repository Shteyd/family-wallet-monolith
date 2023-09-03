package postgres

import (
	"context"
	"monolith/internal/domain"
	"monolith/internal/infrastructure/database"
	"monolith/internal/module/customer/repository/shared"
)

type PostgresAdapter[T shared.CustomerModel] interface {
	Exec(ctx context.Context, sql string, args ...any) error
	Query(ctx context.Context, sql string, args ...any) (T, error)
}

type _PostgresAdapter[T shared.CustomerModel] struct {
	Database domain.Database
}

func NewPostgresAdapter[T shared.CustomerModel](database domain.Database) PostgresAdapter[T] {
	return &_PostgresAdapter[T]{Database: database}
}

func (adapter *_PostgresAdapter[T]) Exec(ctx context.Context, sql string, args ...any) error {
	_, err := adapter.Database.Exec(ctx, sql, args...)
	return err
}

func (adapter *_PostgresAdapter[T]) Query(ctx context.Context, sql string, args ...any) (T, error) {
	rows, err := adapter.Database.Query(ctx, sql, args...)
	if err != nil {
		return T{}, err
	}

	result, err := database.CollectOneRow[T](rows)
	if err != nil {
		return T{}, err
	}

	return *result, nil
}
