package clock

import "time"

// Clock is the injectable time source. Production uses System; tests use Fixed so
// token and OTP expiry are deterministic.
type Clock interface {
	Now() time.Time
}

type systemClock struct{}

func System() Clock { return systemClock{} }

func (systemClock) Now() time.Time { return time.Now() }

// Fixed is a Clock frozen at T, for tests.
type Fixed struct{ T time.Time }

func (f Fixed) Now() time.Time { return f.T }
