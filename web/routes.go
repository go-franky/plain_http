package web

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-franky/plain_http/graphql"
	"github.com/go-franky/plain_http/pq"
	"github.com/gorilla/mux"
)

func (s *Server) routes() {
	router := mux.NewRouter()
	router.HandleFunc("/", s.log(s.rootHandler()))
	router.HandleFunc("/health", s.log(s.health()))
	router.HandleFunc("/graphiql", playground.Handler("GraphQL playground", "/graphql"))
	gql := handler.New(
		graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{}}),
	)
	gql.AddTransport(transport.POST{})
	gql.Use(&debug.Tracer{})
	router.HandleFunc("/graphql", s.log(func(w http.ResponseWriter, r *http.Request) {
		gql.ServeHTTP(w, r)
	})).Methods("POST")

	pq, _ := pq.New()

	pqRoutes := router.PathPrefix("/pq").Subrouter()
	pq.Handlers(pqRoutes)
	router.HandleFunc("/pq/clients/new", pq.AddClient())

	s.Handler = router
}

func (s *Server) rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		s.health()(w, r)
	}
}
