package clock_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/wigata-intech/logres/internal/shared/clock"
)

func TestSystemNow(t *testing.T) {
	before := time.Now()
	got := clock.System().Now()
	after := time.Now()

	assert.False(t, got.Before(before))
	assert.False(t, got.After(after))
}

func TestFixedNow(t *testing.T) {
	want := time.Date(2026, 7, 3, 12, 0, 0, 0, time.UTC)
	c := clock.Fixed{T: want}

	assert.True(t, want.Equal(c.Now()))
	assert.True(t, want.Equal(c.Now())) // stable across calls
}
