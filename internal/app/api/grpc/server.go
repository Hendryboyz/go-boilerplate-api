package grpc

import (
	todogrpc "go-boilerplate-api/pkg/proto/todo/v1"

	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
)

// NewGrpcServer creates and configures a new gRPC server with all services
// This is the main entry point for gRPC server setup
func NewGrpcServer(
	todoHandler *GRPCTodoHandler,
) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(apmgrpc.NewUnaryServerInterceptor()),
	)

	// Register all gRPC services
	todogrpc.RegisterTodoServiceServer(s, todoHandler)

	return s
}
