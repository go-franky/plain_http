package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"moul.io/http2curl"
)

const (
	timeFormat = "2006-01-02T15:04:05.000Z"
)

type Server struct {
	logger  Logger
	Handler http.Handler
	cache   graphql.Cache
}

// New creates and initializes a server
func New(options ...func(s *Server) error) (*Server, error) {
	srv := &Server{
		logger: NoopLogger,
		cache:  nil,
	}

	for _, opt := range options {
		if err := opt(srv); err != nil {
			return &Server{}, err
		}
	}
	srv.routes()

	return srv, nil
}

func SetGraphQLCache(c graphql.Cache) func(s *Server) error {
	return func(s *Server) error {
		s.cache = c
		return nil
	}
}

// WithLogger defines a logger for the server
func WithLogger(l Logger) func(s *Server) error {
	return func(s *Server) error {
		s.logger = l
		return nil
	}
}

type responseWriter struct {
	http.ResponseWriter
}

func (r *responseWriter) Write(b []byte) (int, error) {
	fmt.Println("Written", string(b))
	return r.ResponseWriter.Write(b)
}

func (r *responseWriter) WriteHeader(i int) {
	fmt.Println("Header", i)
	r.ResponseWriter.WriteHeader(i)
}

func (s *Server) log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		command, _ := http2curl.GetCurlCommand(r)
		fmt.Println(command)

		startTime := time.Now()
		h(&responseWriter{w}, r)
		s.logger.Infof("%v %s %s %s", time.Now().UTC().Format(timeFormat), r.Method, r.URL.Path, time.Since(startTime))
	}
}
