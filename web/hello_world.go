package web

import (
	"fmt"
	"net/http"
)

func (s *Server) helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
