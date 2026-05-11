package bootstrap

import (
	internal "github.com/dev-AdiR/cache-management/Internal"
	env "github.com/dev-AdiR/cache-management/Internal/Env"
	types "github.com/dev-AdiR/cache-management/Types"
)

type App struct {
	Client types.Client
}

func Bootstrap(enum int) *App {

	env := env.LoadEnv()

	// TODO: Replace hardcoded Client provider and make it configurable
	registerClient := internal.DbMapper[int(enum)]

	Client := registerClient(env.ClientUrl)
	return &App{
		Client: Client,
	}
}
