package internal

import (
	cache "github.com/dev-AdiR/cache-management/Internal/Cache"
	enums "github.com/dev-AdiR/cache-management/Internal/Enums"
	types "github.com/dev-AdiR/cache-management/Types"
)

const (
	Redis = iota
)

var DbMapper = map[int]ClientFactory{
	int(enums.Redis): cache.RegisterRedisClient,
}

type ClientFactory func(url string) types.Client
