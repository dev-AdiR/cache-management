package main

import (
	enums "github.com/dev-AdiR/cache-management/Enums"
	types "github.com/dev-AdiR/cache-management/Types"
)

type App struct {
	Client types.Client
}

func Bootstrap() *App {

	env := LoadEnv()

	// TODO: Replace hardcoded Client provider and make it configurable
	registerClient := types.DbMapper[int(enums.Redis)]

	Client := registerClient(env.ClientUrl)
	return &App{
		Client: Client,
	}
}
