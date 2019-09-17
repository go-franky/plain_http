package graphql

import (
	"context"
	"runtime"
	"time"

	"github.com/go-franky/plain_http/api"
	"github.com/go-franky/plain_http/version"
)

func (r *queryResolver) Health(ctx context.Context) (*api.Health, error) {
	gitVer := version.GitRevision
	health := &api.Health{
		Alive:     true,
		GitSha:    nil,
		GoVersion: runtime.Version(),
		Timestamp: time.Now(),
	}
	if gitVer != "" {
		health.GitSha = &gitVer
	}

	return health, nil
}
