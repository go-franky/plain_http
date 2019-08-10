package web

import (
	"fmt"
	"net/http"

	"github.com/go-franky/plain_http/version"
)

func (s *Server) health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, fmt.Sprintf(`{"alive":"true","revision":"%s"}`, version.GitRevision))
	}
}
