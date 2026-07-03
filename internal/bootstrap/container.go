package bootstrap

import (
	"context"


	redisclient "github.com/redis/go-redis/v9"

	httpin "github.com/zineb-b/identity-verification-platform-go/internal/adapter/in/http"
	"github.com/zineb-b/identity-verification-platform-go/internal/adapter/out/id"
	clock "github.com/zineb-b/identity-verification-platform-go/internal/adapter/out/time"

	postgres "github.com/zineb-b/identity-verification-platform-go/internal/adapter/out/postgres"
	"github.com/zineb-b/identity-verification-platform-go/internal/adapter/out/security"
	"github.com/zineb-b/identity-verification-platform-go/internal/application/service"
	"github.com/zineb-b/identity-verification-platform-go/internal/config"
	portout "github.com/zineb-b/identity-verification-platform-go/internal/application/port/out"
	redisadapter "github.com/zineb-b/identity-verification-platform-go/internal/adapter/out/redis"


)

// Ce container, cree la cnx avec la base de données, ensuite cree les depandances, à la fin return un object
// avec les handlers

type Container struct {
	HealthHandler *httpin.HealthHandler
	AuthHandler   *httpin.AuthHandler
	TokenManager portout.TokenManager
	RedisClient *redisclient.Client
	Close func()
}

func Build(ctx context.Context, cfg config.Config) (*Container, error) {
	dbPool, err := postgres.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	redisClient, err := redisadapter.NewClient(ctx, cfg.RedisAddr)
	if err != nil {
		dbPool.Close()
		return nil, err
	}

	if err := dbPool.Ping(ctx); err != nil {
		dbPool.Close()
		redisClient.Close()
		return nil, err
	}

	userRepository := postgres.NewUserRepository(dbPool)
	hasher := security.NewBcryptHasher()
	idGenerator := id.NewUUIDGenerator()
	systemClock := clock.NewSystemClock()
	tokenManager := security.NewJWTManager(cfg.JWTSecret)
	

	authService := service.NewAuthService(userRepository, hasher, idGenerator, systemClock, tokenManager)

	return &Container{
		HealthHandler: httpin.NewHealthHandler(dbPool, redisClient),
		AuthHandler:   httpin.NewAuthHandler(authService),
		TokenManager:  tokenManager,
		RedisClient:   redisClient,
		Close: func() {
			dbPool.Close()
		},
	}, nil
}