package web

import (
	"fmt"
	"net/http"
)

func (s *Server) helloWorld() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	}
}
