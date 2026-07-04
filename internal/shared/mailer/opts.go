package mailer

import (
	"net/http"
	"strings"
)

type OptionFunc func(*mailgunMailer)

// WithBaseURL overrides the Mailgun API host, e.g. for the EU region.
func WithBaseURL(baseURL string) OptionFunc {
	return func(m *mailgunMailer) { m.baseURL = strings.TrimSuffix(baseURL, "/") }
}

// WithHTTPClient overrides the client used for the API call.
func WithHTTPClient(c *http.Client) OptionFunc {
	return func(m *mailgunMailer) { m.httpClient = c }
}
