package api

type Health struct {
	Alive     bool    `json:"alive"`
	GitSha    *string `json:"gitSha,omitempty"`
	GoVersion string  `json:"goversion"`
}
