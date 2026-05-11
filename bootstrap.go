package main

import (
	enums "Client-Management/Enums"
	types "Client-Management/Types"
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
