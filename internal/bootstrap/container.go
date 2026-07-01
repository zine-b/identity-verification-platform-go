package bootstrap

import (
	"context"

	httpin "github.com/zineb-b/identity-verification-platform-go/internal/adapter/in/http"
	postgres "github.com/zineb-b/identity-verification-platform-go/internal/adapter/out/postgres"
	"github.com/zineb-b/identity-verification-platform-go/internal/application/service"
	"github.com/zineb-b/identity-verification-platform-go/internal/config"
)

type Container struct {
	HealthHandler *httpin.HealthHandler
	AuthHandler   *httpin.AuthHandler

	Close func()
}

func Build(ctx context.Context, cfg config.Config) (*Container, error) {
	dbPool, err := postgres.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := dbPool.Ping(ctx); err != nil {
		dbPool.Close()
		return nil, err
	}

	userRepository := postgres.NewUserRepository(dbPool)
	authService := service.NewAuthService(userRepository)

	return &Container{
		HealthHandler: httpin.NewHealthHandler(dbPool),
		AuthHandler:   httpin.NewAuthHandler(authService),
		Close: func() {
			dbPool.Close()
		},
	}, nil
}
