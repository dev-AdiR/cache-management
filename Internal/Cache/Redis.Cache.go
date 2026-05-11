package cache

import (
	"context"
	"time"

	types "github.com/dev-AdiR/cache-management/Types"
	"github.com/redis/go-redis/v9"
)

func RegisterRedisClient(redisUrl string) types.Client {
	rdp := redis.NewClient(&redis.Options{
		Addr: redisUrl,
	})

	// rdp.Close()
	return &RedisClient{
		Client: rdp,
	}
}

type RedisClient struct {
	Client *redis.Client
}

func (redis *RedisClient) Set(ctx context.Context, key string, value []byte, expiration time.Duration) error {
	err := redis.Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (redis *RedisClient) Get(ctx context.Context, key string) ([]byte, error) {
	return nil, nil
}

func (redis *RedisClient) Close() error {
	return redis.Client.Close()
}
