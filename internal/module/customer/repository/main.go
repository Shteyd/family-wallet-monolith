package repository

import (
	"monolith/internal/module/customer/adapter/postgres"
	"monolith/internal/module/customer/adapter/redis"
	"monolith/internal/module/customer/core"
)

type _CustomerRepository struct {
	PostgresAdapter postgres.PostgresAdapter
	RedisAdapter    redis.RedisAdapter
}

func NewCustomerRepository(
	postgresAdapter postgres.PostgresAdapter,
	redisAdapter redis.RedisAdapter,
) core.CustomerRepository {
	return &_CustomerRepository{
		PostgresAdapter: postgresAdapter,
		RedisAdapter:    redisAdapter,
	}
}
