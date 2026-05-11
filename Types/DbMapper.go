package types

import enums "Client-Management/Enums"

const (
	Redis = iota
)

var DbMapper = map[int]ClientFactory{
	int(enums.Redis): RegisterRedisClient,
}

type ClientFactory func(url string) Client
