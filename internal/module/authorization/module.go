package authorization

import (
	"monolith/internal/domain"
	"monolith/internal/module/authorization/core"
	"monolith/internal/module/authorization/usecase"
	customerCacheAdapter "monolith/internal/module/customer/adapter/cache"
	customerPostgresAdapter "monolith/internal/module/customer/adapter/postgres"
	customerRepository "monolith/internal/module/customer/repository"
	customerShared "monolith/internal/module/customer/repository/shared"
	passwordCryptoAdapter "monolith/internal/module/password/adapter/cryptohash"
	passwordRepository "monolith/internal/module/password/repository"
	tokenTokenAdapter "monolith/internal/module/token/adapter/token"
	tokenRepository "monolith/internal/module/token/repository"
	"time"
)

type Dependency struct {
	Logger     domain.Logger
	Cache      domain.Cache
	Database   domain.DatabaseManager
	CryptoHash domain.CryptoManager
	Token      domain.TokenManager
	Timeout    time.Duration
}

func NewAuthorizationModule(dependency Dependency) core.AuthorizationUsecase {
	customerCacheAdapter := customerCacheAdapter.NewCacheAdapter(dependency.Cache)
	customerPostgresAdapter := customerPostgresAdapter.NewPostgresManagerAdapter[customerShared.CustomerModel](dependency.Database)
	customerRepository := customerRepository.NewCustomerRepository(
		customerCacheAdapter,
		customerPostgresAdapter,
	)

	passwordCryptoAdapter := passwordCryptoAdapter.NewCryptoHashAdapter(dependency.CryptoHash)
	passwordRepository := passwordRepository.NewPasswordRepository(passwordCryptoAdapter)

	tokenTokenAdapter := tokenTokenAdapter.NewTokenAdapter(dependency.Token)
	tokenRepository := tokenRepository.NewTokenRepository(tokenTokenAdapter)

	return usecase.NewAuthorizationUsecase(
		dependency.Logger,
		customerRepository,
		passwordRepository,
		tokenRepository,
		dependency.Timeout,
	)
}
