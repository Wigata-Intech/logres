package web

import (
	"net/http"

	"github.com/wigata-intech/logres/config"
)

type server struct {
	cfg *config.WebServerConfig
}

func New(opts ...OptionFunc) (http.Handler, error) {
	s := &server{}

	for _, o := range opts {
		o(s)
	}

	mux := http.NewServeMux()

	return mux, nil
}
