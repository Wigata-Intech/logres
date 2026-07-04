package config_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/config"
)

const validJWTSecret = "01234567890123456789012345678901" // 33 bytes, above the 32-byte floor

const (
	envJWTSecret         = "AUTH_JWT_SECRET"
	envPort              = "PORT"
	envDBPath            = "DB_PATH"
	envAccessTTL         = "AUTH_ACCESS_TTL"
	envRefreshTTL        = "AUTH_REFRESH_TTL"
	envArgon2Memory      = "AUTH_ARGON2_MEMORY"
	envArgon2Time        = "AUTH_ARGON2_TIME"
	envArgon2Threads     = "AUTH_ARGON2_THREADS"
	envArgon2Concurrency = "AUTH_ARGON2_CONCURRENCY"
	envOTPTTL            = "OTP_TTL"
	envOTPAttemptMax     = "OTP_ATTEMPT_MAX"
	envOTPLength         = "OTP_LENGTH"
	envMailgunAPIKey     = "MAILGUN_API_KEY"
	envMailgunDomain     = "MAILGUN_DOMAIN"
	envMailgunBaseURL    = "MAILGUN_BASE_URL"
	envMailgunFrom       = "MAILGUN_FROM"
)

// envKeys lists every env var Load reads. Tests clear all of them before each
// case so a variable leaking from the host environment can't taint a case.
var envKeys = []string{
	envJWTSecret, envPort, envDBPath,
	envAccessTTL, envRefreshTTL,
	envArgon2Memory, envArgon2Time, envArgon2Threads, envArgon2Concurrency,
	envOTPTTL, envOTPAttemptMax, envOTPLength,
	envMailgunAPIKey, envMailgunDomain, envMailgunBaseURL, envMailgunFrom,
}

func TestLoad(t *testing.T) {
	type testCase struct {
		name    string
		env     map[string]string
		wantErr bool
		check   func(t *testing.T, cfg *config.Config)
	}

	cases := []testCase{
		{
			name: "valid env fully populated",
			env: map[string]string{
				envJWTSecret:         validJWTSecret,
				envPort:              "9090",
				envDBPath:            "/data/logres.db",
				envAccessTTL:         "5m",
				envRefreshTTL:        "24h",
				envArgon2Memory:      "131072",
				envArgon2Time:        "5",
				envArgon2Threads:     "2",
				envArgon2Concurrency: "8",
				envOTPTTL:            "1m",
				envOTPAttemptMax:     "3",
				envOTPLength:         "6",
				envMailgunAPIKey:     "key-123",
				envMailgunDomain:     "mg.example.com",
				envMailgunBaseURL:    "https://api.eu.mailgun.net",
				envMailgunFrom:       "noreply@example.com",
			},
			check: func(t *testing.T, cfg *config.Config) {
				assert.Equal(t, "9090", cfg.API.Port)
				assert.Equal(t, "/data/logres.db", cfg.DB.Path)
				assert.Equal(t, validJWTSecret, cfg.Auth.JWTSecret)
				assert.Equal(t, 5*time.Minute, cfg.Auth.OAuthToken.AccessTTL)
				assert.Equal(t, 24*time.Hour, cfg.Auth.OAuthToken.RefreshTTL)
				assert.Equal(t, uint32(131072), cfg.Auth.Argon2.Memory)
				assert.Equal(t, uint32(5), cfg.Auth.Argon2.Time)
				assert.Equal(t, uint8(2), cfg.Auth.Argon2.Threads)
				assert.Equal(t, 8, cfg.Auth.Argon2.Concurrency)
				assert.Equal(t, time.Minute, cfg.Auth.OTP.TTL)
				assert.Equal(t, 3, cfg.Auth.OTP.AttemptMax)
				assert.Equal(t, 6, cfg.Auth.OTP.Length)
				assert.Equal(t, "key-123", cfg.Auth.Mailgun.APIKey)
				assert.Equal(t, "mg.example.com", cfg.Auth.Mailgun.Domain)
				assert.Equal(t, "https://api.eu.mailgun.net", cfg.Auth.Mailgun.BaseURL)
				assert.Equal(t, "noreply@example.com", cfg.Auth.Mailgun.From)
			},
		},
		{
			name: "defaults applied when unset",
			env: map[string]string{
				envJWTSecret: validJWTSecret,
			},
			check: func(t *testing.T, cfg *config.Config) {
				assert.Equal(t, "8080", cfg.API.Port)
				assert.Equal(t, "logres.db", cfg.DB.Path)
				assert.Equal(t, 15*time.Minute, cfg.Auth.OAuthToken.AccessTTL)
				assert.Equal(t, 720*time.Hour, cfg.Auth.OAuthToken.RefreshTTL)
				assert.Equal(t, uint32(65536), cfg.Auth.Argon2.Memory)
				assert.Equal(t, uint32(3), cfg.Auth.Argon2.Time)
				assert.Equal(t, uint8(1), cfg.Auth.Argon2.Threads)
				assert.Equal(t, 4, cfg.Auth.Argon2.Concurrency)
				assert.Equal(t, 10*time.Minute, cfg.Auth.OTP.TTL)
				assert.Equal(t, 5, cfg.Auth.OTP.AttemptMax)
				assert.Equal(t, 8, cfg.Auth.OTP.Length)
				assert.Equal(t, "https://api.mailgun.net", cfg.Auth.Mailgun.BaseURL)
				assert.Empty(t, cfg.Auth.Mailgun.APIKey)
				assert.Empty(t, cfg.Auth.Mailgun.Domain)
				assert.Empty(t, cfg.Auth.Mailgun.From)
			},
		},
		{
			name: "missing jwt secret",
			env: map[string]string{
				envJWTSecret: "",
			},
			wantErr: true,
		},
		{
			name: "jwt secret too short",
			env: map[string]string{
				envJWTSecret: "too-short",
			},
			wantErr: true,
		},
		{
			name: "malformed access ttl",
			env: map[string]string{
				envJWTSecret: validJWTSecret,
				envAccessTTL: "not-a-duration",
			},
			wantErr: true,
		},
		{
			name: "malformed argon2 memory",
			env: map[string]string{
				envJWTSecret:    validJWTSecret,
				envArgon2Memory: "not-a-number",
			},
			wantErr: true,
		},
		{
			name: "malformed otp attempt max",
			env: map[string]string{
				envJWTSecret:     validJWTSecret,
				envOTPAttemptMax: "not-a-number",
			},
			wantErr: true,
		},
		{
			name: "argon2 threads out of uint8 range",
			env: map[string]string{
				envJWTSecret:     validJWTSecret,
				envArgon2Threads: "256",
			},
			wantErr: true,
		},
		{
			name: "negative access ttl rejected",
			env: map[string]string{
				envJWTSecret: validJWTSecret,
				envAccessTTL: "-5m",
			},
			wantErr: true,
		},
		{
			name: "refresh ttl not exceeding access rejected",
			env: map[string]string{
				envJWTSecret:  validJWTSecret,
				envAccessTTL:  "1h",
				envRefreshTTL: "30m",
			},
			wantErr: true,
		},
		{
			name: "argon2 memory below floor rejected",
			env: map[string]string{
				envJWTSecret:    validJWTSecret,
				envArgon2Memory: "1024",
			},
			wantErr: true,
		},
		{
			name: "otp length below minimum rejected",
			env: map[string]string{
				envJWTSecret: validJWTSecret,
				envOTPLength: "4",
			},
			wantErr: true,
		},
		{
			name: "otp length above maximum rejected",
			env: map[string]string{
				envJWTSecret: validJWTSecret,
				envOTPLength: "12",
			},
			wantErr: true,
		},
		{
			name: "partial mailgun config rejected",
			env: map[string]string{
				envJWTSecret:     validJWTSecret,
				envMailgunAPIKey: "key-only",
			},
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for _, key := range envKeys {
				t.Setenv(key, "")
			}
			for k, v := range tc.env {
				t.Setenv(k, v)
			}

			cfg, err := config.Load()
			if tc.wantErr {
				require.Error(t, err)
				assert.Nil(t, cfg)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, cfg)
			tc.check(t, cfg)
		})
	}
}

func TestLoadWeb(t *testing.T) {
	type testCase struct {
		name string
		port string
		want string
	}

	cases := []testCase{
		{name: "default port when unset", want: "8081"},
		{name: "port override", port: "9091", want: "9091"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(envPort, tc.port)

			cfg, err := config.LoadWeb()
			require.NoError(t, err)
			require.NotNil(t, cfg)
			assert.Equal(t, tc.want, cfg.Port)
		})
	}
}

func TestLoadWebAdmin(t *testing.T) {
	type testCase struct {
		name string
		port string
		want string
	}

	cases := []testCase{
		{name: "default port when unset", want: "8082"},
		{name: "port override", port: "9092", want: "9092"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(envPort, tc.port)

			cfg, err := config.LoadWebAdmin()
			require.NoError(t, err)
			require.NotNil(t, cfg)
			assert.Equal(t, tc.want, cfg.Port)
		})
	}
}
