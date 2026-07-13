package bootstrap

import (
	"context"
	"log/slog"

	redisclient "github.com/redis/go-redis/v9"
	loggeradapter "github.com/zineb-b/identity-verification-platform-go/internal/adapter/out/logger"

	httpin "github.com/zineb-b/identity-verification-platform-go/internal/adapter/in/http"
	"github.com/zineb-b/identity-verification-platform-go/internal/adapter/out/id"
	clock "github.com/zineb-b/identity-verification-platform-go/internal/adapter/out/time"

	postgres "github.com/zineb-b/identity-verification-platform-go/internal/adapter/out/postgres"
	redisadapter "github.com/zineb-b/identity-verification-platform-go/internal/adapter/out/redis"
	"github.com/zineb-b/identity-verification-platform-go/internal/adapter/out/security"
	portout "github.com/zineb-b/identity-verification-platform-go/internal/application/port/out"
	"github.com/zineb-b/identity-verification-platform-go/internal/application/service"
	"github.com/zineb-b/identity-verification-platform-go/internal/config"
)

// Ce container, cree la cnx avec la base de données, ensuite cree les depandances, à la fin return un object
// avec les handlers

type Container struct {
	HealthHandler *httpin.HealthHandler
	AuthHandler   *httpin.AuthHandler
	TokenManager  portout.TokenManager
	Logger        *slog.Logger
	RedisClient   *redisclient.Client
	Close         func()
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
	appLogger := loggeradapter.NewJSONLogger(cfg.Env)

	authService := service.NewAuthService(userRepository, hasher, idGenerator, systemClock, tokenManager)

	return &Container{
		HealthHandler: httpin.NewHealthHandler(dbPool, redisClient),
		AuthHandler:   httpin.NewAuthHandler(authService),
		TokenManager:  tokenManager,
		RedisClient:   redisClient,
		Logger:        appLogger,
		Close: func() {
			dbPool.Close()
		},
	}, nil
}
