// Package config loads Logres runtime configuration from the environment.
// Load() is the single entry point: it applies defaults, then fails fast on
// anything malformed or missing so a bad deploy never reaches the point of
// serving traffic.
package config

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/wigata-intech/logres/internal/shared/envx"
)

// API

type APIServerConfig struct {
	Port string
}

// Admin authentication

// AdminOAuthTokenConfig holds the admin access/refresh token lifetimes.
type AdminOAuthTokenConfig struct {
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

// AdminOTPConfig holds the password-reset OTP policy.
type AdminOTPConfig struct {
	TTL        time.Duration
	AttemptMax int
	Length     int
}

// Argon2Config holds the password-hashing cost parameters (see
// internal/shared/hasher.Params) plus the in-flight-hash concurrency cap.
type Argon2Config struct {
	Memory      uint32
	Time        uint32
	Threads     uint8
	Concurrency int
}

// MailgunConfig holds the Mailgun credentials. A zero value means Mailgun is
// not configured; callers fall back to a Noop mailer rather than failing.
type MailgunConfig struct {
	APIKey  string
	Domain  string
	BaseURL string
	From    string
}

// AuthConfig groups everything the admin-auth service depends on.
type AuthConfig struct {
	JWTSecret  string
	OAuthToken AdminOAuthTokenConfig
	OTP        AdminOTPConfig
	Argon2     Argon2Config
	Mailgun    MailgunConfig
}

// DBConfig controls where the SQLite file lives.
type DBConfig struct {
	Path string
}

// Web

type WebServerConfig struct {
	Port string
}

// Web Admin

type WebAdminServerConfig struct {
	Port string
}

// Config is the root configuration for the api surface, loaded once at
// startup via Load.
type Config struct {
	API  APIServerConfig
	DB   DBConfig
	Auth AuthConfig
}

const (
	defaultPort           = "8080"
	defaultWebPort        = "8081"
	defaultWebAdminPort   = "8082"
	defaultDBPath         = "logres.db"
	defaultAccessTTL      = 15 * time.Minute
	defaultRefreshTTL     = 720 * time.Hour
	defaultArgon2Memory   = 64 * 1024
	defaultArgon2Time     = 3
	defaultArgon2Threads  = 1
	defaultArgon2Concur   = 4
	defaultOTPTTL         = 10 * time.Minute
	defaultOTPAttemptMax  = 5
	defaultOTPLength      = 8
	defaultMailgunBaseURL = "https://api.mailgun.net"

	// minJWTSecretLen mirrors internal/shared/token's RFC 8725 §3.5 floor: an
	// HMAC key at least as long as the HS256 output (256 bits).
	minJWTSecretLen = 32

	// Sanity floors — reject nonsense config, not a security recommendation.
	minArgon2MemoryKiB = 8 * 1024 // 8 MiB
	minOTPLength       = 6
	maxOTPLength       = 10
)

// Load reads configuration from the environment, applying defaults for
// anything unset. It fails fast on a missing/weak JWT secret or a malformed
// value — never a partially-valid Config.
func Load() (*Config, error) {
	jwtSecret := os.Getenv("AUTH_JWT_SECRET")
	if len(jwtSecret) < minJWTSecretLen {
		return nil, fmt.Errorf("config.Load: AUTH_JWT_SECRET must be set and at least %d bytes", minJWTSecretLen)
	}

	accessTTL, err := envx.Duration("AUTH_ACCESS_TTL", defaultAccessTTL)
	if err != nil {
		return nil, fmt.Errorf("config.Load: %w", err)
	}
	refreshTTL, err := envx.Duration("AUTH_REFRESH_TTL", defaultRefreshTTL)
	if err != nil {
		return nil, fmt.Errorf("config.Load: %w", err)
	}

	argon2Memory, err := envx.Uint32("AUTH_ARGON2_MEMORY", defaultArgon2Memory)
	if err != nil {
		return nil, fmt.Errorf("config.Load: %w", err)
	}
	argon2Time, err := envx.Uint32("AUTH_ARGON2_TIME", defaultArgon2Time)
	if err != nil {
		return nil, fmt.Errorf("config.Load: %w", err)
	}
	argon2Threads, err := envx.Uint32("AUTH_ARGON2_THREADS", defaultArgon2Threads)
	if err != nil {
		return nil, fmt.Errorf("config.Load: %w", err)
	}
	if argon2Threads > math.MaxUint8 {
		return nil, fmt.Errorf("config.Load: AUTH_ARGON2_THREADS must fit in a byte, got %d", argon2Threads)
	}
	argon2Concurrency, err := envx.Int("AUTH_ARGON2_CONCURRENCY", defaultArgon2Concur)
	if err != nil {
		return nil, fmt.Errorf("config.Load: %w", err)
	}

	otpTTL, err := envx.Duration("OTP_TTL", defaultOTPTTL)
	if err != nil {
		return nil, fmt.Errorf("config.Load: %w", err)
	}
	otpAttemptMax, err := envx.Int("OTP_ATTEMPT_MAX", defaultOTPAttemptMax)
	if err != nil {
		return nil, fmt.Errorf("config.Load: %w", err)
	}
	otpLength, err := envx.Int("OTP_LENGTH", defaultOTPLength)
	if err != nil {
		return nil, fmt.Errorf("config.Load: %w", err)
	}

	cfg := &Config{
		API: APIServerConfig{
			Port: envx.Get("PORT", defaultPort),
		},
		DB: DBConfig{
			Path: envx.Get("DB_PATH", defaultDBPath),
		},
		Auth: AuthConfig{
			JWTSecret: jwtSecret,
			OAuthToken: AdminOAuthTokenConfig{
				AccessTTL:  accessTTL,
				RefreshTTL: refreshTTL,
			},
			OTP: AdminOTPConfig{
				TTL:        otpTTL,
				AttemptMax: otpAttemptMax,
				Length:     otpLength,
			},
			Argon2: Argon2Config{
				Memory:      argon2Memory,
				Time:        argon2Time,
				Threads:     uint8(argon2Threads), // range-checked above
				Concurrency: argon2Concurrency,
			},
			Mailgun: MailgunConfig{
				APIKey:  os.Getenv("MAILGUN_API_KEY"),
				Domain:  os.Getenv("MAILGUN_DOMAIN"),
				BaseURL: envx.Get("MAILGUN_BASE_URL", defaultMailgunBaseURL),
				From:    os.Getenv("MAILGUN_FROM"),
			},
		},
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("config.Load: %w", err)
	}
	return cfg, nil
}

// validate rejects parseable-but-nonsensical values so a bad deploy fails at
// startup rather than at the first request. Type/format errors are already
// handled during parsing; this covers semantic ranges and cross-field rules,
// delegated to each sub-config so no single function grows unwieldy.
func (c *Config) validate() error {
	for _, check := range []func() error{
		c.Auth.OAuthToken.validate,
		c.Auth.Argon2.validate,
		c.Auth.OTP.validate,
		c.Auth.Mailgun.validate,
	} {
		if err := check(); err != nil {
			return err
		}
	}
	return nil
}

func (c AdminOAuthTokenConfig) validate() error {
	if c.AccessTTL <= 0 {
		return fmt.Errorf("AUTH_ACCESS_TTL must be positive, got %s", c.AccessTTL)
	}
	if c.RefreshTTL <= c.AccessTTL {
		return fmt.Errorf("AUTH_REFRESH_TTL (%s) must exceed AUTH_ACCESS_TTL (%s)", c.RefreshTTL, c.AccessTTL)
	}
	return nil
}

func (c Argon2Config) validate() error {
	if c.Memory < minArgon2MemoryKiB {
		return fmt.Errorf("AUTH_ARGON2_MEMORY must be >= %d KiB, got %d", minArgon2MemoryKiB, c.Memory)
	}
	if c.Time < 1 {
		return fmt.Errorf("AUTH_ARGON2_TIME must be >= 1")
	}
	if c.Threads < 1 {
		return fmt.Errorf("AUTH_ARGON2_THREADS must be >= 1")
	}
	if c.Concurrency < 1 {
		return fmt.Errorf("AUTH_ARGON2_CONCURRENCY must be >= 1, got %d", c.Concurrency)
	}
	return nil
}

func (c AdminOTPConfig) validate() error {
	if c.TTL <= 0 {
		return fmt.Errorf("OTP_TTL must be positive, got %s", c.TTL)
	}
	if c.AttemptMax < 1 {
		return fmt.Errorf("OTP_ATTEMPT_MAX must be >= 1, got %d", c.AttemptMax)
	}
	if c.Length < minOTPLength || c.Length > maxOTPLength {
		return fmt.Errorf("OTP_LENGTH must be between %d and %d, got %d", minOTPLength, maxOTPLength, c.Length)
	}
	return nil
}

func (c MailgunConfig) validate() error {
	anySet := c.APIKey != "" || c.Domain != "" || c.From != ""
	allSet := c.APIKey != "" && c.Domain != "" && c.From != ""
	if anySet && !allSet {
		return fmt.Errorf("MAILGUN_API_KEY, MAILGUN_DOMAIN and MAILGUN_FROM must be set together or all left empty")
	}
	return nil
}

// LoadWeb reads the public storefront server's configuration from the
// environment. Unlike Load, it has no auth dependency, so a missing
// AUTH_JWT_SECRET never blocks the web binary from starting.
func LoadWeb() (*WebServerConfig, error) {
	return &WebServerConfig{
		Port: envx.Get("PORT", defaultWebPort),
	}, nil
}

// LoadWebAdmin reads the admin panel server's configuration from the
// environment. Like LoadWeb, it is auth-independent.
func LoadWebAdmin() (*WebAdminServerConfig, error) {
	return &WebAdminServerConfig{
		Port: envx.Get("PORT", defaultWebAdminPort),
	}, nil
}
