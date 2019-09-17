package api

import "time"

type Health struct {
	Alive     bool      `json:"alive"`
	GitSha    *string   `json:"gitSha,omitempty"`
	GoVersion string    `json:"goversion"`
	Timestamp time.Time `json:"timestamp"`
}
