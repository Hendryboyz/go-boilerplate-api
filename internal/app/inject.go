package app

import (
	"go-boilerplate-api/internal/app/api"
	"go-boilerplate-api/internal/app/repository"
	"go-boilerplate-api/internal/app/todo"

	"github.com/google/wire"
)

var repositoryBindings = wire.NewSet(
	wire.Bind(new(todo.TodoRepository), new(*repository.GormTodoRepository)),
	repository.ProviderSet,
)

var ApiProviderSet = wire.NewSet(
	repositoryBindings,
	todo.NewTodoService,
	api.ProviderSet,
)
