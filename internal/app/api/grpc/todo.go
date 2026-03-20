package grpc

import (
	"go-boilerplate-api/internal/app/todo"
	todov1 "go-boilerplate-api/pkg/proto/todo/v1"
)

type GRPCTodoHandler struct {
	todov1.UnimplementedTodoServiceServer
	todoService todo.TodoService
}

func NewGRPCTodoHandler(
	todoService todo.TodoService,
) *GRPCTodoHandler {
	return &GRPCTodoHandler{
		todoService: todoService,
	}
}
