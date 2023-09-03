package cache

import (
	"fmt"
	"monolith/internal/domain"
	"monolith/internal/module/customer/core"
	"time"
)

const (
	defaultCustomerExpiration = 24 * time.Hour
)

func getCustomerKey(entity core.Customer) string {
	return fmt.Sprintf("customer-%d", entity.Id)
}

type CacheAdapter interface {
	Set(core.Customer)
	Get(core.Customer) (core.Customer, error)
	Del(core.Customer)
}

type _CacheAdapter struct {
	Cache domain.Cache
}

func NewCacheAdapter(cache domain.Cache) CacheAdapter {
	return &_CacheAdapter{Cache: cache}
}

func (adapter *_CacheAdapter) Del(entity core.Customer) {
	adapter.Cache.Del(
		getCustomerKey(entity),
	)
}

func (adapter *_CacheAdapter) Get(entity core.Customer) (core.Customer, error) {
	customer, err := adapter.Cache.Get(getCustomerKey(entity))
	if err != nil {
		return core.Customer{}, err
	}

	entity, ok := customer.(core.Customer)
	if !ok {
		return core.Customer{}, domain.ErrorInvalidType
	}

	return entity, nil
}

func (adapter *_CacheAdapter) Set(entity core.Customer) {
	adapter.Cache.Set(
		getCustomerKey(entity),
		entity,
		defaultCustomerExpiration,
	)
}
