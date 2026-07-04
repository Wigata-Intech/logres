package config

import "time"

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

// Web

type WebServerConfig struct {
	Port string
}

// Web Admin

type WebAdminServerConfig struct {
	Port string
}
