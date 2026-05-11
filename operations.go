package main

import (
	"context"
	"time"

	types "github.com/dev-AdiR/cache-management/Types"
)

type Cache struct {
	client types.Client
}

func (cache *Cache) Add(ctx context.Context, key string, value []byte, exp time.Duration) error {

	err := cache.client.Set(ctx, key, value, exp)

	return err
}

func (cache *Cache) Get(ctx context.Context, key string) ([]byte, error) {

	data, err := cache.client.Get(ctx, key)

	if err != nil {
		return nil, err
	}

	return data, nil
}
