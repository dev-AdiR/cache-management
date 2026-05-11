package types

import enums "github.com/dev-AdiR/cache-management/Enums"

const (
	Redis = iota
)

var DbMapper = map[int]ClientFactory{
	int(enums.Redis): RegisterRedisClient,
}

type ClientFactory func(url string) Client
