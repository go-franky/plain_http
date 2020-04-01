package pq

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type gqlClientOperations struct{}

type gqlDigestOperation struct {
	digest string
	body   string
	name   string
}

type memoryRepo struct {
	clients map[string]string
}

func (m *memoryRepo) Clients() []GraphQLClient, error {
	for
	return []GraphQLClient{}, nil
}


type clientRepo interface {
	AddClient(c GQLClient) client
	Clients() []GraphQLClient, error
}

// var _ PersistedQueryRepo = &memoryRepo{}

type PersistedQueryRepo interface{}

type (
	Server struct {
		Handler http.Handler
		repo    PersistedQueryRepo
	}
)

//func New(dataStore PersistedQueryRepo) (*http.ServeMux, error) {
func New() (*Server, error) {
	fmt.Println("HERE")

	router := http.NewServeMux()
	router.HandleFunc("/add_client", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HERE")
	})

	return &Server{
		Handler: router,
	}, nil
}

func (s *Server) Router() http.Handler {
	return s.Handler
}

func (s *Server) AddClient() error {
}

func (s *Server) Handlers(m *mux.Router) {
	m.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode()
	}).Methods("GET")

	m.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request) {
		err := s.AddClient(NewClient())
	}).Methods("POST")
}
