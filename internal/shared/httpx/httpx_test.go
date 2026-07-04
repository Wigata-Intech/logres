package httpx_test

import (
	"context"
	"io"
	"log/slog"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/internal/shared/httpx"
)

func TestServe(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	type testCase struct {
		name    string
		addr    func(t *testing.T) string
		ctx     func() (context.Context, context.CancelFunc)
		wantErr bool
	}

	cases := []testCase{
		{
			name: "already cancelled context shuts down cleanly",
			addr: func(*testing.T) string { return "127.0.0.1:0" },
			ctx: func() (context.Context, context.CancelFunc) {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx, cancel
			},
		},
		{
			name: "context deadline shuts down cleanly",
			addr: func(*testing.T) string { return "127.0.0.1:0" },
			ctx: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 5*time.Millisecond)
			},
		},
		{
			name: "listen error surfaces before context is done",
			addr: func(t *testing.T) string {
				t.Helper()
				var lc net.ListenConfig
				l, err := lc.Listen(context.Background(), "tcp", "127.0.0.1:0")
				require.NoError(t, err)
				t.Cleanup(func() { _ = l.Close() })
				return l.Addr().String()
			},
			ctx: func() (context.Context, context.CancelFunc) {
				return context.WithCancel(context.Background())
			},
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := tc.ctx()
			defer cancel()

			err := httpx.Serve(ctx, tc.addr(t), http.NewServeMux(), logger)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
