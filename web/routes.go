package web

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-franky/plain_http/graphql"
)

func (s *Server) routes() {
	router := http.NewServeMux()
	router.HandleFunc("/", s.log(s.rootHandler()))
	router.HandleFunc("/health", s.log(s.health()))
	router.HandleFunc("/graphiql", handler.Playground("GraphQL playground", "/graphql"))
	router.HandleFunc("/graphql", s.log(handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{}}))))
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
