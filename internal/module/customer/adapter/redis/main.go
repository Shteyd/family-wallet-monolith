package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisAdapter interface {
	Set(ctx context.Context, key string, value interface{}) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, key string) *redis.IntCmd
}

type _RedisAdapter struct {
	Client *redis.Client
}

func NewRedisAdapter(client *redis.Client, objectTtl time.Duration) RedisAdapter {
	return &_RedisAdapter{
		Client: client,
	}
}

func (adapter *_RedisAdapter) getObjectTTL() time.Duration {
	return 1 * time.Hour
}

func (adapter *_RedisAdapter) Get(ctx context.Context, key string) *redis.StringCmd {
	return adapter.Client.Get(ctx, key)
}

func (adapter *_RedisAdapter) Set(ctx context.Context, key string, value interface{}) *redis.StatusCmd {
	return adapter.Client.Set(ctx, key, value, adapter.getObjectTTL())
}

func (adapter *_RedisAdapter) Del(ctx context.Context, key string) *redis.IntCmd {
	return adapter.Client.Del(ctx, key)
}
