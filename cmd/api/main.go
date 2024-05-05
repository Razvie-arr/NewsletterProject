package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"newsletterProject/config"
	"newsletterProject/mailer"
	"newsletterProject/pkg/authenticator"
	"newsletterProject/repository"

	"newsletterProject/service"
	"newsletterProject/transport/api"
	"newsletterProject/transport/util"

	httpx "go.strv.io/net/http"
)

var version = "v1.0.1"

func main() {
	ctx := context.Background()
	cfg := config.MustLoadConfig()
	util.SetServerLogLevel(slog.LevelInfo)

	database, err := setupDatabase(ctx, cfg)
	if err != nil {
		slog.Error("initializing database", slog.Any("error", err))
	}
	repo, err := repository.New(database)
	if err != nil {
		slog.Error("initializing repository", slog.Any("error", err))
	}
	resendMailer := mailer.NewResendMailer(cfg.ResendApiKey)

	controller, err := setupController(
		cfg,
		repo,
		resendMailer,
	)
	if err != nil {
		slog.Error("initializing controller", slog.Any("error", err))
	}

	addr := fmt.Sprintf(":%d", cfg.Port)
	// Initialize the server config.
	serverConfig := httpx.ServerConfig{
		Addr:    addr,
		Handler: controller,
		Hooks: httpx.ServerHooks{
			BeforeShutdown: []httpx.ServerHookFunc{
				func(_ context.Context) {
					database.Close()
				},
			},
		},
		Limits: nil,
		Logger: util.NewServerLogger("httpx.Server"),
	}
	server := httpx.NewServer(&serverConfig)

	slog.Info("starting server", slog.Int("port", cfg.Port))
	if err := server.Run(ctx); err != nil {
		slog.Error("server failed", slog.Any("error", err))
	}
}

func setupDatabase(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	// Initialize the database connection pool.
	pool, err := pgxpool.New(
		ctx,
		cfg.DatabaseURL,
	)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func setupController(
	cfg config.Config,
	repository service.Repository,
	mailer mailer.Mailer,
) (*api.Controller, error) {
	// Initialize the service.
	svc, err := service.NewService(repository, mailer)
	if err != nil {
		return nil, fmt.Errorf("initializing editor service: %w", err)
	}

	JWTAuthenticator := authenticator.NewJWTAuthenticator(cfg.SupabaseAuthSecret)

	// Initialize the controller.
	controller, err := api.NewController(
		JWTAuthenticator,
		svc,
		version,
	)
	if err != nil {
		return nil, fmt.Errorf("initializing controller: %w", err)
	}

	return controller, nil
}
