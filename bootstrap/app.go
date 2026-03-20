package bootstrap

import (
	"context"
	"fmt"
	"go-boilerplate-api/global"
	"go-boilerplate-api/internal/pkg/log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewHttpServer(router *gin.Engine) *http.Server {
	port := global.App.Config.Server.HttpPort
	startEndpoint := fmt.Sprintf(":%d", port)
	InitializeValidator()
	return &http.Server{
		Addr:    startEndpoint,
		Handler: router.Handler(),
	}
}

type App struct {
	httpService  *http.Server
	grpcServer   *grpc.Server
	grpcListener net.Listener
}

func NewGrpcListener() (net.Listener, error) {
	grpcPort := global.App.Config.Server.GrpcPort
	if grpcPort == 0 {
		grpcPort = 9000
	}
	return net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
}

func NewApiApp(
	httpService *http.Server,
	grpcServer *grpc.Server,
	grpcListener net.Listener,
) *App {
	return &App{
		httpService:  httpService,
		grpcServer:   grpcServer,
		grpcListener: grpcListener,
	}
}

func (a *App) Run() error {
	go func() {
		log.Info(fmt.Sprintf("http server starting listen at %s", a.httpService.Addr))
		if err := a.httpService.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(fmt.Sprintf("server failed to listen at %s\n", err))
		}
	}()

	go func() {
		log.Info(fmt.Sprintf("gRPC server starting listen at %s", a.grpcListener.Addr()))
		if err := a.grpcServer.Serve(a.grpcListener); err != nil && err != grpc.ErrServerStopped {
			log.Fatal(fmt.Sprintf("gRPC server failed to serve: %v\n", err))
		}
	}()

	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	// Create a channel to signal when shutdown is complete
	done := make(chan struct{})
	var httpErr, grpcErr error

	// Shutdown HTTP server
	go func() {
		if a.httpService != nil {
			httpErr = a.httpService.Shutdown(ctx)
		}
		done <- struct{}{}
	}()

	// Shutdown gRPC server
	go func() {
		if a.grpcServer != nil {
			a.grpcServer.GracefulStop()
		}
		done <- struct{}{}
	}()

	// Wait for both servers to shut down
	for i := 0; i < 2; i++ {
		<-done
	}

	// Return the first error that occurred, if any
	if httpErr != nil {
		return fmt.Errorf("HTTP server shutdown error: %v", httpErr)
	}
	if grpcErr != nil {
		return fmt.Errorf("gRPC server shutdown error: %v", grpcErr)
	}

	return nil
}
