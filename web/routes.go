package web

import (
	"net/http"
)

func (s *Server) routes() {
	router := http.NewServeMux()
	router.HandleFunc("/", s.log(s.rootHandler()))
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
