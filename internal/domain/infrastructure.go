package domain

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type (
	LoggerArgs map[string]any

	Logger interface {
		Debug(msg string, args LoggerArgs)
		Error(err error, args LoggerArgs)
		Warn(msg string, args LoggerArgs)
		Info(msg string, args LoggerArgs)
	}
)

type (
	Cache interface {
		Set(string, any, time.Duration)
		Get(string) (any, error)
		Del(string)
	}
)

type (
	CryptoManager interface {
		Encrypt(string) (string, error)
	}
)

type (
	DatabaseManager interface {
		Begin(context.Context) (DatabaseManager, error)
		Commit(context.Context) error
		Rollback(context.Context) error

		GetConnect() Database
	}

	Database interface {
		Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
		Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
		QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
		SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	}
)
