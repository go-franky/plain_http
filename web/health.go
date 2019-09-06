package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"github.com/go-franky/plain_http/version"
)

func (s *Server) health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		health := struct {
			Alive     bool   `json:"alive"`
			Revision  string `json:"revision,omitempty"`
			GoVersion string `json:"goversion,omitempty"`
		}{
			Alive:     true,
			Revision:  version.GitRevision,
			GoVersion: runtime.Version(),
		}
		data, err := json.Marshal(health)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, string(data))
	}
}
