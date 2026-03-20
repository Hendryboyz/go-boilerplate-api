package api

import (
	"go-boilerplate-api/internal/app/api/grpc"
	"go-boilerplate-api/internal/app/api/http"
	httpHandlers "go-boilerplate-api/internal/app/api/http/handlers"

	"github.com/google/wire"
)

var grpcHandlerProvider = wire.NewSet(
	grpc.NewGRPCTodoHandler,
)

var httpHandlerProviderSet = wire.NewSet(
	httpHandlers.NewTodoController,
)

var grpcServerProvider = wire.NewSet(grpc.NewGrpcServer, grpcHandlerProvider)
var httpServerProvider = wire.NewSet(http.NewApiRouter, httpHandlerProviderSet)

var ProviderSet = wire.NewSet(httpServerProvider, grpcServerProvider)
