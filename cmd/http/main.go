package main

import (
	"context"
	"monolith/config"
	"monolith/internal/infrastructure/cache"
	"monolith/internal/infrastructure/crypto"
	"monolith/internal/infrastructure/database"
	"monolith/internal/infrastructure/delivery/http"
	"monolith/internal/infrastructure/jwt"
	"monolith/internal/infrastructure/logger"
	"monolith/internal/module/authorization"
	"monolith/pkg/pgxpool"
	"monolith/pkg/zerolog"
	"os"
	"os/signal"
	"time"
)

const (
	defaultUsecaseTimeout = 5 * time.Second
)

func main() {
	// signals
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// init config
	config := config.New(".")

	// init logger
	zerolog := zerolog.NewStdout(config.IsDebug)
	logger := logger.New(zerolog)

	// init datasources
	cache := cache.NewCache(time.Hour, time.Hour)
	pgxpool := pgxpool.NewDatabase(ctx, config.DatabaseDsn)
	defer pgxpool.Close()

	database := database.NewDatabaseManager(pgxpool, nil)

	// crypto
	cryptoManager := crypto.NewCryptoManager(config.PasswordSalt)

	// jwt
	jwtManager := jwt.NewTokenManager(config.TokenSalt)

	// init modules
	authModule := authorization.NewAuthorizationModule(authorization.Dependency{
		Logger:     logger,
		Cache:      cache,
		Database:   database,
		CryptoHash: cryptoManager,
		Token:      jwtManager,
		Timeout:    defaultUsecaseTimeout,
	})

	server := http.NewHttpServer(authModule, config.HttpPort)
	server.Run()
}
