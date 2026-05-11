package types

import enums "github.com/dev-AdiR/cache-management/Internal/Enums"

const (
	Redis = iota
)

var DbMapper = map[int]ClientFactory{
	int(enums.Redis): RegisterRedisClient,
}

type ClientFactory func(url string) Client
