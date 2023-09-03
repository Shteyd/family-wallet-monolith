package repository

import (
	"monolith/internal/module/customer/adapter/cache"
	"monolith/internal/module/customer/adapter/postgres"
	"monolith/internal/module/customer/core"
)

type _CustomerRepository struct {
	CacheAdapter    cache.CacheAdapter
	PostgresAdapter postgres.PostgresAdapter
}

func NewCustomerRepository(
	cacheAdapter cache.CacheAdapter,
	postgresAdapter postgres.PostgresAdapter,
) core.CustomerRepository {
	return &_CustomerRepository{
		CacheAdapter:    cacheAdapter,
		PostgresAdapter: postgresAdapter,
	}
}
