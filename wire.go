//go:build wireinject
// +build wireinject

package main

import (
	"go-boilerplate-api/bootstrap"
	"go-boilerplate-api/internal/app"
	"go-boilerplate-api/internal/pkg"

	"github.com/google/wire"
)

func wireApp() (*bootstrap.App, error) {
	panic(
		wire.Build(
			pkg.DataProviderSet,
			app.ApiProviderSet,
			bootstrap.ProviderSet,
		),
	)
}
