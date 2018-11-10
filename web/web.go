package web

import (
	"fmt"
	"net/http"
	"time"
)

const (
	timeFormat = "2006-01-02T15:04:05.000Z"
)

type Server struct {
	Port   int
	Logger Logger
}

// Start begins the HTTP server
func (s *Server) Start() error {
	s.Logger.Infof("Web server now listing on port %d", s.Port)
	http.HandleFunc("/", s.Log(s.helloWorld))
	return http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil)
}

// New creates and initializes a server
func New(options ...func(s *Server) error) (*Server, error) {
	srv := &Server{
		Port: 8080,
	}

	for _, opt := range options {
		if err := opt(srv); err != nil {
			return &Server{}, err
		}
	}

	return srv, nil
}

// WithLogger defines a logger for the server
func WithLogger(l Logger) func(s *Server) error {
	return func(s *Server) error {
		s.Logger = l
		return nil
	}
}

// WithPort set the port on which the server should listen on
func WithPort(l int) func(s *Server) error {
	return func(s *Server) error {
		s.Port = l
		return nil
	}
}

func (s *Server) Log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		h(w, r)
		s.Logger.Infof("%v %s %s %s", time.Now().UTC().Format(timeFormat), r.Method, r.URL.Path, time.Since(startTime))
	}
}