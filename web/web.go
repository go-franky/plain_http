package web

import (
	"net/http"
	"time"
)

const (
	timeFormat = "2006-01-02T15:04:05.000Z"
)

// Option configures a Server.
type Option func(s *Server) error

type Server struct {
	logger  Logger
	Handler http.Handler
}

// New creates and initializes a server
func New(options ...Option) (*Server, error) {
	srv := &Server{
		logger: NoopLogger,
	}

	for _, opt := range options {
		if err := opt(srv); err != nil {
			return &Server{}, err
		}
	}
	srv.routes()

	return srv, nil
}

// WithLogger defines a logger for the server
func WithLogger(l Logger) Option {
	return func(s *Server) error {
		s.logger = l
		return nil
	}
}

func (s *Server) log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		h(w, r)
		s.logger.Infof("%v %s %s %s", time.Now().UTC().Format(timeFormat), r.Method, r.URL.Path, time.Since(startTime))
	}
}
