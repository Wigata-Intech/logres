// Package mailer sends transactional email (e.g. password-reset OTP codes) via
// Mailgun's HTTP API. No SDK: the API surface is one POST, so stdlib
// net/http keeps the dependency count flat.
package mailer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// defaultBaseURL targets Mailgun's US region. Operators in the EU/ID pass
// WithBaseURL("https://api.eu.mailgun.net") for data residency.
const defaultBaseURL = "https://api.mailgun.net"

const defaultTimeout = 10 * time.Second

// Message is a single email send request.
type Message struct {
	To      string
	Subject string
	Text    string
	HTML    string
}

// Mailer sends Message. Interface keeps SES/SMTP swappable later without
// touching callers.
type Mailer interface {
	Send(ctx context.Context, msg Message) error
}

type mailgunMailer struct {
	apiKey     string
	domain     string
	from       string
	baseURL    string
	httpClient *http.Client
}

type Option func(*mailgunMailer)

// WithBaseURL overrides the Mailgun API host, e.g. for the EU region.
func WithBaseURL(baseURL string) Option {
	return func(m *mailgunMailer) { m.baseURL = strings.TrimSuffix(baseURL, "/") }
}

// WithHTTPClient overrides the client used for the API call.
func WithHTTPClient(c *http.Client) Option {
	return func(m *mailgunMailer) { m.httpClient = c }
}

// New builds a Mailgun-backed Mailer. apiKey, domain, and from are required.
func New(apiKey, domain, from string, opts ...Option) (Mailer, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("mailer.New: apiKey is required")
	}
	if domain == "" {
		return nil, fmt.Errorf("mailer.New: domain is required")
	}
	if from == "" {
		return nil, fmt.Errorf("mailer.New: from is required")
	}

	m := &mailgunMailer{
		apiKey:     apiKey,
		domain:     domain,
		from:       from,
		baseURL:    defaultBaseURL,
		httpClient: &http.Client{Timeout: defaultTimeout},
	}
	for _, opt := range opts {
		opt(m)
	}
	return m, nil
}

func (m *mailgunMailer) Send(ctx context.Context, msg Message) error {
	form := url.Values{
		"from":    {m.from},
		"to":      {msg.To},
		"subject": {msg.Subject},
		"text":    {msg.Text},
	}
	if msg.HTML != "" {
		form.Set("html", msg.HTML)
	}

	endpoint := fmt.Sprintf("%s/v3/%s/messages", m.baseURL, m.domain)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("mailer.Send: %w", err)
	}
	req.SetBasicAuth("api", m.apiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("mailer.Send: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// Status code only: the response body may echo back request details.
		return fmt.Errorf("mailer.Send: %w: status %d", errNonSuccess, resp.StatusCode)
	}
	return nil
}

var errNonSuccess = errors.New("mailgun returned non-2xx response")

// Noop discards messages. Used for local/dev where no Mailgun credentials
// are configured; a logging wrapper can be layered on later if needed.
type Noop struct{}

func (Noop) Send(context.Context, Message) error { return nil }
