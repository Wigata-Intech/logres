// Command logres-api runs the Logres HTTP API server.
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/wigata-intech/logres/config"

	// Shared Packages
	"github.com/wigata-intech/logres/internal/shared/cli"
	"github.com/wigata-intech/logres/internal/shared/envx"
	"github.com/wigata-intech/logres/internal/shared/hasher"
	"github.com/wigata-intech/logres/internal/shared/httpx"
	"github.com/wigata-intech/logres/internal/shared/mailer"
	"github.com/wigata-intech/logres/internal/shared/sqlite"
	"github.com/wigata-intech/logres/internal/shared/token"

	// Repositories
	adminPasswordResetRepository "github.com/wigata-intech/logres/internal/api/repository/adminPasswordReset"
	adminRefreshTokenRepository "github.com/wigata-intech/logres/internal/api/repository/adminRefreshToken"
	adminUserRepository "github.com/wigata-intech/logres/internal/api/repository/adminUser"
	cartRepository "github.com/wigata-intech/logres/internal/api/repository/cart"
	categoryRepository "github.com/wigata-intech/logres/internal/api/repository/category"
	fileRepository "github.com/wigata-intech/logres/internal/api/repository/file"
	orderRepository "github.com/wigata-intech/logres/internal/api/repository/order"
	productRepository "github.com/wigata-intech/logres/internal/api/repository/product"
	tagRepository "github.com/wigata-intech/logres/internal/api/repository/tag"

	// Services
	adminFileService "github.com/wigata-intech/logres/internal/api/service/adminFile"
	adminOrderService "github.com/wigata-intech/logres/internal/api/service/adminOrder"
	adminProductService "github.com/wigata-intech/logres/internal/api/service/adminProduct"
	adminUserService "github.com/wigata-intech/logres/internal/api/service/adminUser"
	cartService "github.com/wigata-intech/logres/internal/api/service/cart"
	orderService "github.com/wigata-intech/logres/internal/api/service/order"
	productService "github.com/wigata-intech/logres/internal/api/service/product"

	// Handler
	"github.com/wigata-intech/logres/internal/api"
	"github.com/wigata-intech/logres/internal/api/handler/admin"
	"github.com/wigata-intech/logres/internal/api/handler/cart"
	"github.com/wigata-intech/logres/internal/api/handler/order"
	"github.com/wigata-intech/logres/internal/api/handler/product"
)

// buildAPIHandler wires the full DI graph for the api surface: repositories
// over db, shared auth primitives from cfg, services over those, handlers
// over those services, and finally the routed http.Handler.
func buildAPIHandler(logger *slog.Logger, db *sql.DB, cfg *config.Config) (http.Handler, error) {
	// Shared Packages
	h := hasher.New(
		hasher.WithParams(hasher.Params{
			Memory:  cfg.Auth.Argon2.Memory,
			Time:    cfg.Auth.Argon2.Time,
			Threads: cfg.Auth.Argon2.Threads,
			SaltLen: 16,
			KeyLen:  32,
		}),
		hasher.WithConcurrency(cfg.Auth.Argon2.Concurrency),
	)

	issuer, err := token.New([]byte(cfg.Auth.JWTSecret))
	if err != nil {
		return nil, fmt.Errorf("buildAPIHandler: token.New: %w", err)
	}

	var mail mailer.Mailer
	mg := cfg.Auth.Mailgun
	if mg.APIKey != "" && mg.Domain != "" && mg.From != "" {
		mail, err = mailer.New(mg.APIKey, mg.Domain, mg.From, mailer.WithBaseURL(mg.BaseURL))
		if err != nil {
			return nil, fmt.Errorf("buildAPIHandler: mailer.New: %w", err)
		}
	} else {
		logger.Warn("mailgun not configured; using noop mailer")
		mail = mailer.Noop{}
	}

	// Repository
	adminUserRepo := adminUserRepository.New(db, logger)
	adminPasswordResetRepo := adminPasswordResetRepository.New(db, logger)
	adminRefreshTokenRepo := adminRefreshTokenRepository.New(db, logger)
	cartRepo := cartRepository.New(db, logger)
	categoryRepo := categoryRepository.New(db, logger)
	fileRepo := fileRepository.New(db, logger)
	orderRepo := orderRepository.New(db, logger)
	productRepo := productRepository.New(db, logger)
	tagRepo := tagRepository.New(db, logger)

	// Service
	adminUserSvc := adminUserService.New(
		adminUserService.WithLogger(logger),
		adminUserService.WithHasher(h),
		adminUserService.WithTokenIssuer(issuer),
		adminUserService.WithMailer(mail),
		adminUserService.WithAdminUserRepository(adminUserRepo),
		adminUserService.WithAdminPasswordResetRepository(adminPasswordResetRepo),
		adminUserService.WithAdminRefreshTokenRepository(adminRefreshTokenRepo),
		adminUserService.WithOAuthTokenConfig(cfg.Auth.OAuthToken),
		adminUserService.WithOTPConfig(cfg.Auth.OTP),
	)
	adminFileSvc := adminFileService.New(logger, fileRepo)
	adminProductSvc := adminProductService.New(logger, categoryRepo, fileRepo, productRepo, tagRepo)
	adminOrderSvc := adminOrderService.New(logger, cartRepo, orderRepo)
	cartSvc := cartService.New(logger, cartRepo, productRepo)
	orderSvc := orderService.New(logger, cartRepo, orderRepo, productRepo)
	productSvc := productService.New(logger, productRepo)

	// Handler
	adminHandler := admin.New(adminUserSvc, adminFileSvc, adminProductSvc, adminOrderSvc)
	cartHandler := cart.New(cartSvc)
	orderHandler := order.New(orderSvc)
	productHandler := product.New(productSvc)

	return api.New(
		api.WithAdminHandler(adminHandler),
		api.WithCartHandler(cartHandler),
		api.WithOrderHandler(orderHandler),
		api.WithProductHandler(productHandler),
	)
}

func main() {
	logger := slog.Default()
	if err := run(context.Background(), logger); err != nil {
		logger.Error("api", "error", err)
		os.Exit(1)
	}
}

// run holds everything that needs its deferred cleanup to actually execute —
// main just maps its error to an exit code.
func run(ctx context.Context, logger *slog.Logger) error {
	if err := envx.LoadFile(".env", logger); err != nil {
		return fmt.Errorf("env: %w", err)
	}

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}

	db, err := sqlite.Open(ctx, sqlite.Config{Path: cfg.DB.Path})
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	defer func() { _ = db.Close() }()

	handler, err := buildAPIHandler(logger, db, cfg)
	if err != nil {
		return fmt.Errorf("build handler: %w", err)
	}

	logger.Info("logres-api starting", "version", cli.Version)

	if err := httpx.Serve(ctx, ":"+cfg.API.Port, handler, logger); err != nil {
		return fmt.Errorf("serve: %w", err)
	}
	return nil
}
