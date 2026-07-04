package token

import "time"

type OptionFunc func(*hs256Issuer)

func WithIssuer(iss string) OptionFunc   { return func(i *hs256Issuer) { i.issuer = iss } }
func WithAudience(aud string) OptionFunc { return func(i *hs256Issuer) { i.audience = aud } }

// WithClock overrides the time source (default time.Now) for deterministic tests.
func WithClock(now func() time.Time) OptionFunc {
	return func(i *hs256Issuer) {
		if now != nil {
			i.now = now
		}
	}
}
