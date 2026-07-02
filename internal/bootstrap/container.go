package bootstrap

import (
	"context"

	httpin "github.com/zineb-b/identity-verification-platform-go/internal/adapter/in/http"
	postgres "github.com/zineb-b/identity-verification-platform-go/internal/adapter/out/postgres"
	"github.com/zineb-b/identity-verification-platform-go/internal/application/service"
	"github.com/zineb-b/identity-verification-platform-go/internal/config"
	"github.com/zineb-b/identity-verification-platform-go/internal/adapter/out/security"
	"github.com/zineb-b/identity-verification-platform-go/internal/adapter/out/id"

)

// Ce container, cree la cnx avec la base de données, ensuite cree les depandances, à la fin return un object 
// avec les handlers

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
	hasher := security.NewBcryptHasher()
	idGenerator := id.NewUUIDGenerator()
	authService := service.NewAuthService(userRepository, hasher, idGenerator)

	return &Container{
		HealthHandler: httpin.NewHealthHandler(dbPool),
		AuthHandler:   httpin.NewAuthHandler(authService),
		Close: func() {
			dbPool.Close()
		},
	}, nil
}
