package postgres

import (
	"context"
	"monolith/internal/domain"
)

type PostgresAdapter interface {
	domain.DatabaseManager
}

type _PostgresAdapter struct {
	DatabaseManager domain.DatabaseManager
}

func NewPostgresAdapter(databaseManager domain.DatabaseManager) PostgresAdapter {
	return &_PostgresAdapter{
		DatabaseManager: databaseManager,
	}
}

func (adapter *_PostgresAdapter) Begin(ctx context.Context) (domain.DatabaseManager, error) {
	return adapter.DatabaseManager.Begin(ctx)
}

func (adapter *_PostgresAdapter) Commit(ctx context.Context) error {
	return adapter.DatabaseManager.Commit(ctx)
}

func (adapter *_PostgresAdapter) Rollback(ctx context.Context) error {
	return adapter.DatabaseManager.Rollback(ctx)
}

func (adapter *_PostgresAdapter) GetConnect() domain.Database {
	return adapter.DatabaseManager.GetConnect()
}
