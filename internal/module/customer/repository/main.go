package repository

import (
	"monolith/internal/module/customer/adapter/cache"
	"monolith/internal/module/customer/adapter/postgres"
	"monolith/internal/module/customer/core"
	"monolith/internal/module/customer/repository/shared"
)

type _CustomerRepository struct {
	CacheAdapter    cache.CacheAdapter
	PostgresAdapter postgres.PostgresManagerAdapter[shared.CustomerModel]
}

func NewCustomerRepository(
	cacheAdapter cache.CacheAdapter,
	postgresAdapter postgres.PostgresManagerAdapter[shared.CustomerModel],
) core.CustomerRepository {
	return &_CustomerRepository{
		CacheAdapter:    cacheAdapter,
		PostgresAdapter: postgresAdapter,
	}
}
