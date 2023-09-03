package customer

import (
	"monolith/internal/domain"
	"monolith/internal/module/customer/adapter/cache"
	"monolith/internal/module/customer/adapter/postgres"
	"monolith/internal/module/customer/core"
	"monolith/internal/module/customer/repository"
	"monolith/internal/module/customer/usecase"
	"time"
)

type Dependency struct {
	Logger   domain.Logger
	Cache    domain.Cache
	Database domain.DatabaseManager
	Timeout  time.Duration
}

func NewCustomerModule(dependency Dependency) core.CustomerUsecase {
	customerCacheAdapter := cache.NewCacheAdapter(dependency.Cache)
	customerPostgresAdapter := postgres.NewPostgresAdapter(dependency.Database)

	customerRepository := repository.NewCustomerRepository(
		customerCacheAdapter,
		customerPostgresAdapter,
	)

	return usecase.NewCustomerUsecase(
		dependency.Logger,
		customerRepository,
		dependency.Timeout,
	)
}
