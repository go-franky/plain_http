package pq

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/99designs/gqlgen/graphql"
)

type clientSecret string

type queryHash string

var registeredQueries map[clientSecret]map[queryHash]string

type GQLClient struct {
	Name   string
	Secret string
}

type client interface{}

func NewClient(name string) client {
	hex, _ := randomHex(32)
	return &GQLClient{
		Name:   name,
		Secret: hex,
	}
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (c *GQLClient) AddOperation(name graphql.RawParams) error {
	return nil
}
