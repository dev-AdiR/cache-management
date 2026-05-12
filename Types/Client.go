package types

import (
	"context"
	"time"
)

type Client interface {
	Set(ctx context.Context, key string, value []byte, expiration time.Duration) error
	Get(ctx context.Context, key string) ([]byte, error)
	Close() error
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) any
}
