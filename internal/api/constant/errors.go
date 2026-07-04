package constant

import "errors"

// Business errors surfaced by the api service/repository layers. The handler
// maps each to an HTTP status and client message (RFC 9457) at the boundary;
// the mapping lives with the handler, not here.
var (
	ErrAdminAlreadyExists = errors.New("admin account already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid or expired refresh token")
	ErrInvalidReset       = errors.New("invalid or expired reset code")
)
