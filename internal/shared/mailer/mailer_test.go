package mailer_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/internal/shared/mailer"
)

const (
	sampleKey     = "key-123"
	sampleSubject = "Your OTP"
	sampleDomain  = "mg.example.com"
	sampleText    = "code: 12345678"
	sampleFrom    = "noreply@example.com"
	sampleTo      = "user@example.com"
)

func TestNew(t *testing.T) {
	type testCase struct {
		name    string
		apiKey  string
		domain  string
		from    string
		wantErr bool
	}

	cases := []testCase{
		{name: "valid", apiKey: sampleKey, domain: sampleDomain, from: sampleFrom},
		{name: "missing apiKey", apiKey: "", domain: sampleDomain, from: sampleFrom, wantErr: true},
		{name: "missing domain", apiKey: sampleKey, domain: "", from: sampleFrom, wantErr: true},
		{name: "missing from", apiKey: sampleKey, domain: sampleDomain, from: "", wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m, err := mailer.New(tc.apiKey, tc.domain, tc.from)
			if tc.wantErr {
				require.Error(t, err)
				assert.Nil(t, m)
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, m)
		})
	}
}

func TestMailgunMailerSend(t *testing.T) {
	type testCase struct {
		name       string
		msg        mailer.Message
		statusCode int
		wantErr    bool
	}

	cases := []testCase{
		{
			name:       "success with html",
			msg:        mailer.Message{To: sampleTo, Subject: sampleSubject, Text: sampleText, HTML: "<p>code: 12345678</p>"},
			statusCode: http.StatusOK,
		},
		{
			name:       "success without html",
			msg:        mailer.Message{To: sampleTo, Subject: sampleSubject, Text: sampleText},
			statusCode: http.StatusOK,
		},
		{
			name:       "unauthorized",
			msg:        mailer.Message{To: sampleTo, Subject: sampleSubject, Text: sampleText},
			statusCode: http.StatusUnauthorized,
			wantErr:    true,
		},
		{
			name:       "server error",
			msg:        mailer.Message{To: sampleTo, Subject: sampleSubject, Text: sampleText},
			statusCode: http.StatusInternalServerError,
			wantErr:    true,
		},
	}

	const apiKey = "test-api-key"
	const domain = sampleDomain
	const from = sampleFrom

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var gotMethod, gotPath string
			var gotUser, gotPass string
			var gotHasAuth bool
			var gotForm map[string][]string

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				gotUser, gotPass, gotHasAuth = r.BasicAuth()

				require.NoError(t, r.ParseForm())
				gotForm = map[string][]string(r.PostForm)

				w.WriteHeader(tc.statusCode)
			}))
			defer server.Close()

			m, err := mailer.New(apiKey, domain, from, mailer.WithBaseURL(server.URL))
			require.NoError(t, err)

			err = m.Send(context.Background(), tc.msg)

			assert.Equal(t, http.MethodPost, gotMethod)
			assert.Equal(t, "/v3/"+domain+"/messages", gotPath)
			assert.True(t, gotHasAuth)
			assert.Equal(t, "api", gotUser)
			assert.Equal(t, apiKey, gotPass)
			assert.Equal(t, []string{from}, gotForm["from"])
			assert.Equal(t, []string{tc.msg.To}, gotForm["to"])
			assert.Equal(t, []string{tc.msg.Subject}, gotForm["subject"])
			assert.Equal(t, []string{tc.msg.Text}, gotForm["text"])
			if tc.msg.HTML != "" {
				assert.Equal(t, []string{tc.msg.HTML}, gotForm["html"])
			} else {
				assert.NotContains(t, gotForm, "html")
			}

			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMailgunMailerSendInvalidRequest(t *testing.T) {
	// A control character in the base URL makes http.NewRequestWithContext fail
	// before any network call happens.
	m, err := mailer.New("key", sampleDomain, sampleFrom, mailer.WithBaseURL("http://\x7f"))
	require.NoError(t, err)

	err = m.Send(context.Background(), mailer.Message{To: sampleTo, Subject: "s", Text: "t"})
	require.Error(t, err)
}

func TestMailgunMailerSendServerUnreachable(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	server.Close() // unreachable: connection refused

	m, err := mailer.New("key", sampleDomain, sampleFrom, mailer.WithBaseURL(server.URL))
	require.NoError(t, err)

	err = m.Send(context.Background(), mailer.Message{To: sampleTo, Subject: "s", Text: "t"})
	require.Error(t, err)
}

func TestMailgunMailerSendWithCustomHTTPClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &http.Client{}
	m, err := mailer.New("key", sampleDomain, sampleFrom, mailer.WithBaseURL(server.URL), mailer.WithHTTPClient(client))
	require.NoError(t, err)

	err = m.Send(context.Background(), mailer.Message{To: sampleTo, Subject: "s", Text: "t"})
	require.NoError(t, err)
}

func TestNoopSend(t *testing.T) {
	var n mailer.Noop
	err := n.Send(context.Background(), mailer.Message{To: sampleTo, Subject: "s", Text: "t"})
	assert.NoError(t, err)
}
