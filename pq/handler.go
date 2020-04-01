package pq

import(
	"net/http"
	"fmt"
)

type gqlClientOperations struct {}

type gqlDigestOperation struct {
	digest string
	body string
	name string
}

type memoryRepo struct {
	clients map[string]GQLClient
}

// var _ PersistedQueryRepo = &memoryRepo{}

type PersistedQueryRepo interface {}

type(
	Server struct {
		Handler http.Handler
		repo PersistedQueryRepo
	}
)

//func New(dataStore PersistedQueryRepo) (*http.ServeMux, error) {
func New() (*Server, error) {
	fmt.Println("HERE")

	router := http.NewServeMux()
	router.HandleFunc("/add_client", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("HERE")
	})

	return &Server{
		Handler: router,
	}, nil
}

func (s *Server) Router() http.Handler {
	return s.Handler
}

func (s *Server) AddClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
